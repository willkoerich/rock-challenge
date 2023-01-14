package main

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	handler "github.com/willkoerich/rock-challenge/cmd/api/handler"
	"github.com/willkoerich/rock-challenge/cmd/api/registry"
	"github.com/willkoerich/rock-challenge/cmd/api/response"
	"log"
	"net/http"
	"os"
)

const AccountDescription = "access_account"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	databaseDriver := registry.NewDatabaseDriver()
	passwordGenerator := registry.NewPasswordGenerator()

	accountRepository := registry.NewAccountRepository(databaseDriver)

	accountHandler := handler.NewAccountHandler(registry.NewAccountController(databaseDriver, passwordGenerator))
	loginHandler := handler.NewLoginHandler(registry.NewLoginController(databaseDriver, passwordGenerator))
	transferHandler := handler.NewTransferHandler(registry.NewTransferController(databaseDriver, accountRepository))

	r := mapRoutes(accountHandler, loginHandler, transferHandler)

	logrus.Info("Challenge App Starting...")
	err = http.ListenAndServe(":"+port, r)
	if err != nil {
		return
	}
}

func mapRoutes(accountHandler handler.AccountHandler, loginHandler handler.LoginHandler, transferHandler handler.TransferHandler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/accounts", accountHandler.Create).Methods("POST")
	r.HandleFunc("/accounts", accountHandler.Get).Methods("GET")
	r.HandleFunc("/accounts/{id}", accountHandler.GetByID).Methods("GET")
	r.HandleFunc("/accounts/{id}/balance", accountHandler.GetAccountBalance).Methods("GET")

	r.HandleFunc("/login", loginHandler.Authenticate).Methods("POST")

	r.HandleFunc("/transfers", transferHandler.Get).Methods("GET")
	r.HandleFunc("/transfers", func(writer http.ResponseWriter, request *http.Request) {
		accountID, err := loginHandler.VerifyToken(writer, request)
		if err != nil {
			response.CreateHandlerResponse(writer, http.StatusBadRequest, " acccess_token not accepted", &err)
			return
		}
		ctx := context.WithValue(request.Context(), AccountDescription, accountID)
		updatedRequest := request.WithContext(ctx)
		transferHandler.Process(writer, updatedRequest)
	}).Methods("POST")

	r.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte("pong"))
	}).Methods("GET")
	return r
}
