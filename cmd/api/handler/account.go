package handlers

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/willkoerich/rock-challenge/internal/domain"
	"net/http"
	"strconv"
	"strings"
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
		logrus.Error(" unable to decode submitted account information body. Error: " + err.Error())
		writer.WriteHeader(http.StatusBadRequest)
		_, err := writer.Write([]byte(" unable to decode submitted account information body. Error: " + err.Error()))
		if err != nil {
			return
		}
		return
	}
	createdAccount, err := handler.controller.Create(request.Context(), decodedBody)
	if err != nil {
		logrus.Error(" error to create account")
		writer.WriteHeader(http.StatusInternalServerError)
		_, err := writer.Write([]byte(err.Error()))
		if err != nil {
			return
		}
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(writer).Encode(createdAccount)
	if err != nil {
		return
	}

}

func (handler AccountHandler) GetByID(writer http.ResponseWriter, request *http.Request) {
	accountID := strings.TrimPrefix(request.URL.Path, "/accounts/")
	numericalAccountId, err := strconv.Atoi(accountID)
	if err != nil {
		logrus.Error(" invalid numerical id submitted. Error: " + err.Error())
		writer.WriteHeader(http.StatusBadRequest)
		_, err := writer.Write([]byte(" invalid numerical id submitted. Error: " + err.Error()))
		if err != nil {
			return
		}
		return
	}
	retrievedAccount, err := handler.controller.GetByID(request.Context(), numericalAccountId)
	if err != nil {
		logrus.Errorf(" error to retrieve account by submitted id %v", numericalAccountId)
		writer.WriteHeader(http.StatusInternalServerError)
		_, err := writer.Write([]byte(err.Error()))
		if err != nil {
			return
		}
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	err = json.NewEncoder(writer).Encode(retrievedAccount)
	if err != nil {
		return
	}
}

func (handler AccountHandler) Get(writer http.ResponseWriter, request *http.Request) {
	accounts, err := handler.controller.GetAll(request.Context())
	if err != nil {
		logrus.Error(" error to retrieve accounts")
		writer.WriteHeader(http.StatusInternalServerError)
		_, err := writer.Write([]byte(err.Error()))
		if err != nil {
			return
		}
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	err = json.NewEncoder(writer).Encode(accounts)
	if err != nil {
		return
	}
}
