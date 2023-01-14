package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/willkoerich/rock-challenge/internal/domain"
	domainMocks "github.com/willkoerich/rock-challenge/internal/mocks/domain"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type (
	InvalidBody struct{}
)

func TestCreateAccountHandlerSuccessful(t *testing.T) {

	controller := new(domainMocks.AccountController)
	controller.
		On("Create", mock.Anything, mock.Anything).
		Return(domain.Account{}, nil)

	responseRecorder, request := getContext(domain.Account{}, "/accounts")
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(responseRecorder.Result().Body)
	NewAccountHandler(controller).Create(responseRecorder, request)

	assert.Equal(t, http.StatusCreated, responseRecorder.Code)
}

func TestGetByIDAccountHandlerSuccessful(t *testing.T) {

	controller := new(domainMocks.AccountController)
	controller.
		On("GetByID", mock.Anything, mock.Anything).
		Return(domain.Account{}, nil)

	responseRecorder, request := getContext(domain.Account{}, "/accounts/ ")
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(responseRecorder.Result().Body)
	NewAccountHandler(controller).GetByID(responseRecorder, request)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
}

func TestGetAccountsHandlerSuccessful(t *testing.T) {

	controller := new(domainMocks.AccountController)
	controller.
		On("GetAll", mock.Anything, mock.Anything).
		Return([]domain.Account{}, nil)

	responseRecorder, request := getContext(domain.Account{}, "/accounts")
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(responseRecorder.Result().Body)
	NewAccountHandler(controller).Get(responseRecorder, request)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
}

func getContext(body interface{}, path string) (*httptest.ResponseRecorder, *http.Request) {
	responseRecorder := httptest.NewRecorder()
	if body != nil {
		jsonBody, _ := json.Marshal(body)
		requestReader := bytes.NewReader(jsonBody)
		request := httptest.NewRequest("POST", path, requestReader)
		request.Header.Add("Context-Type", "application/json")
		request = request.WithContext(context.WithValue(request.Context(), AccountDescription, domain.AuthenticationResponse{AccountID: 2}))
		return responseRecorder, request
	}
	return responseRecorder, nil
}
