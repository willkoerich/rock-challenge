package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/go-http-utils/headers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/willkoerich/rock-challenge/cmd/api/response"
	"github.com/willkoerich/rock-challenge/internal/domain"
	"net/http"
	"strconv"
)

const (
	InvalidSubmittedAccountMessage   = "unable to decode submitted account information body."
	ErrorCreateAccountMessage        = "error to create account"
	InvalidSubmittedAccountIdMessage = "invalid numerical id submitted."
	ErrorRetrieveAccountByIdMessage  = "error to retrieve account by submitted id %v"
	ErrorToRetrieveAccountsMessage   = "error to retrieve accounts"
)

type (
	AccountHandler struct {
		controller domain.AccountController
	}
)

func NewAccountHandler(controller domain.AccountController) AccountHandler {
	return AccountHandler{
		controller: controller,
	}
}

func (handler AccountHandler) Create(writer http.ResponseWriter, request *http.Request) {
	var decodedBody domain.CreateRequest
	if err := json.NewDecoder(request.Body).Decode(&decodedBody); err != nil {
		logrus.Error(InvalidSubmittedAccountMessage, err.Error())
		response.CreateHandlerResponse(writer, http.StatusBadRequest,
			InvalidSubmittedAccountMessage, &err)
		return
	}
	createdAccount, err := handler.controller.Create(request.Context(), decodedBody)
	if err != nil {
		logrus.Error(ErrorCreateAccountMessage)
		response.CreateHandlerResponse(writer, http.StatusInternalServerError, ErrorCreateAccountMessage, &err)
		return
	}
	writer.Header().Add(headers.ContentType, "application/json")
	writer.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(writer).Encode(createdAccount)
	if err != nil {
		return
	}

}

func (handler AccountHandler) GetAccountBalance(writer http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(mux.Vars(request)["id"])
	if err != nil {
		logrus.Error(InvalidSubmittedAccountIdMessage, err.Error())
		response.CreateHandlerResponse(writer, http.StatusBadRequest,
			InvalidSubmittedAccountIdMessage, &err)
		return
	}
	retrievedAccount, err := handler.controller.GetByID(request.Context(), id)
	if err != nil {
		logrus.Errorf(ErrorRetrieveAccountByIdMessage, id)
		response.CreateHandlerResponse(writer, http.StatusInternalServerError,
			fmt.Sprintf(ErrorRetrieveAccountByIdMessage, id), &err)
		return
	}
	writer.Header().Add(headers.ContentType, "application/json")
	writer.WriteHeader(http.StatusOK)
	err = json.NewEncoder(writer).Encode(domain.BalanceResponse{Balance: retrievedAccount.Balance})
	if err != nil {
		return
	}
}

func (handler AccountHandler) GetByID(writer http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(mux.Vars(request)["id"])
	if err != nil {
		logrus.Error(InvalidSubmittedAccountIdMessage, err.Error())
		response.CreateHandlerResponse(writer, http.StatusBadRequest,
			InvalidSubmittedAccountIdMessage, &err)
		return
	}
	retrievedAccount, err := handler.controller.GetByID(request.Context(), id)
	if err != nil {
		logrus.Errorf(ErrorRetrieveAccountByIdMessage, id)
		response.CreateHandlerResponse(writer, http.StatusInternalServerError,
			fmt.Sprintf(ErrorRetrieveAccountByIdMessage, id), &err)
		return
	}
	writer.Header().Add(headers.ContentType, "application/json")
	writer.WriteHeader(http.StatusOK)
	err = json.NewEncoder(writer).Encode(retrievedAccount)
	if err != nil {
		return
	}
}

func (handler AccountHandler) Get(writer http.ResponseWriter, request *http.Request) {
	accounts, err := handler.controller.GetAll(request.Context())
	if err != nil {
		logrus.Error(ErrorToRetrieveAccountsMessage)
		response.CreateHandlerResponse(writer, http.StatusInternalServerError,
			ErrorToRetrieveAccountsMessage, &err)
		return
	}
	writer.Header().Add(headers.ContentType, "application/json")
	writer.WriteHeader(http.StatusOK)
	err = json.NewEncoder(writer).Encode(accounts)
	if err != nil {
		return
	}
}
