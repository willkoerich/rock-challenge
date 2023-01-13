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

func TestProcessTransferHandlerSuccessful(t *testing.T) {

	controller := new(domainMocks.TransferController)
	controller.
		On("Save", mock.Anything, mock.Anything).
		Return(domain.Transfer{}, nil)

	responseRecorder, request := getContext(domain.Account{}, "/transfers")
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(responseRecorder.Result().Body)
	NewTransferHandler(controller).Process(responseRecorder, request)

	assert.Equal(t, http.StatusCreated, responseRecorder.Code)
}

func TestGetTransfersHandlerSuccessful(t *testing.T) {

	controller := new(domainMocks.TransferController)
	controller.
		On("GetAll", mock.Anything, mock.Anything).
		Return([]domain.Transfer{}, nil)

	responseRecorder, request := getContext(domain.Account{}, "/transfers")
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(responseRecorder.Result().Body)
	NewTransferHandler(controller).Get(responseRecorder, request)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
}

/*func TestLoginAuthenticateHandlerFailureWhenBodyIsInvalid(t *testing.T) {

	controller := new(domainMocks.TransferController)
	controller.
		On("Authenticate", mock.Anything, mock.Anything).
		Return(domain.Transfer{}, nil)

	responseRecorder, request := getContext(InvalidBody{})
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(responseRecorder.Result().Body)
	NewTransferHandler(controller).Process(responseRecorder, request)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
}*/
