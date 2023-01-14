package handlers

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/willkoerich/rock-challenge/internal/domain"
	domainMocks "github.com/willkoerich/rock-challenge/internal/mocks/domain"
	"io"
	"net/http"
	"testing"
)

func TestLoginAuthenticateHandlerSuccessful(t *testing.T) {

	controller := new(domainMocks.LoginController)
	controller.
		On("Authenticate", mock.Anything, mock.Anything).
		Return(domain.AuthenticationResponse{}, nil)

	responseRecorder, request := getContext(domain.Account{}, "/")
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(responseRecorder.Result().Body)
	NewLoginHandler(controller).Authenticate(responseRecorder, request)

	assert.Equal(t, http.StatusCreated, responseRecorder.Code)
}
