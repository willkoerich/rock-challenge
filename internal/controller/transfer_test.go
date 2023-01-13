package controller

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/willkoerich/rock-challenge/internal/domain"
	domainMock "github.com/willkoerich/rock-challenge/internal/mocks/domain"
	"testing"
)

func TestController_CreateTransferSuccessfully(t *testing.T) {
	repository := new(domainMock.TransferRepository)
	repository.
		On("Save",
			mock.Anything, mock.Anything).
		Return(domain.Transfer{}, nil)

	transfer, err := NewTransferController(repository).Save(context.Background(), domain.Transfer{})

	assert.Equal(t, domain.Transfer{}, transfer)
	assert.Nil(t, err)
}

func TestController_CreateTransferFailure(t *testing.T) {
	repository := new(domainMock.TransferRepository)
	repository.
		On("Save",
			mock.Anything, mock.Anything).
		Return(domain.Transfer{}, errors.New(" error"))

	transfer, err := NewTransferController(repository).Save(context.Background(), domain.Transfer{})

	assert.Equal(t, domain.Transfer{}, transfer)
	assert.NotNil(t, err)
}

func TestController_GetTransfersSuccessfully(t *testing.T) {
	repository := new(domainMock.TransferRepository)
	repository.
		On("GetAll",
			mock.Anything).
		Return([]domain.Transfer{}, nil)

	createdAccount, err := NewTransferController(repository).GetAll(context.Background())

	assert.Equal(t, []domain.Transfer{}, createdAccount)
	assert.Equal(t, nil, err)
}

func TestController_GetTransfersFailure(t *testing.T) {
	repository := new(domainMock.TransferRepository)
	repository.
		On("GetAll",
			mock.Anything).
		Return([]domain.Transfer{}, errors.New(" failure to retrieve accounts"))

	createdAccount, err := NewTransferController(repository).GetAll(context.Background())

	assert.Equal(t, []domain.Transfer{}, createdAccount)
	assert.NotNil(t, err)
}
