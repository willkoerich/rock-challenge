package controller

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/willkoerich/rock-challenge/internal/domain"
	domainMock "github.com/willkoerich/rock-challenge/internal/mocks/domain"
	mocks "github.com/willkoerich/rock-challenge/internal/mocks/plataform/database"
	"testing"
)

func TestController_CreateTransferSuccessfully(t *testing.T) {

	transaction := new(mocks.Transaction)
	transaction.
		On("Commit").
		Return(nil)

	repository := new(domainMock.TransferRepository)
	repository.
		On("Process",
			mock.Anything, mock.Anything, mock.Anything).
		Return(domain.Transfer{}, nil)
	repository.
		On("BeginTransaction",
			mock.Anything).
		Return(transaction, nil)

	accountRepository := new(domainMock.AccountRepository)
	accountRepository.
		On("Update",
			mock.Anything, mock.Anything, mock.Anything).
		Return(nil).Twice()
	accountRepository.
		On("GetByID",
			mock.Anything, mock.Anything).
		Return(domain.Account{}, nil).Twice()

	transfer, err := NewTransferController(repository, accountRepository).
		Process(context.Background(), domain.Transfer{})

	assert.Equal(t, domain.Transfer{}, transfer)
	assert.Nil(t, err)
}

func TestController_CreateTransferWhenOriginHasNoFounds(t *testing.T) {
	repository := new(domainMock.TransferRepository)
	repository.
		On("Process",
			mock.Anything, mock.Anything).
		Return(domain.Transfer{}, nil)

	accountRepository := new(domainMock.AccountRepository)
	accountRepository.
		On("Update",
			mock.Anything, mock.Anything).
		Return(nil).Twice()
	accountRepository.
		On("GetByID",
			mock.Anything, mock.Anything).
		Return(domain.Account{Balance: 10}, nil).Once()
	accountRepository.
		On("GetByID",
			mock.Anything, mock.Anything).
		Return(domain.Account{Balance: 20}, nil).Once()

	_, err := NewTransferController(repository, accountRepository).
		Process(context.Background(), domain.Transfer{Amount: 200})

	assert.NotNil(t, err)
}

func TestController_CreateTransferFailure(t *testing.T) {

	transaction := new(mocks.Transaction)
	transaction.
		On("Commit").
		Return(nil)
	transaction.
		On("Rollback").
		Return(nil)

	repository := new(domainMock.TransferRepository)
	repository.
		On("Process",
			mock.Anything, mock.Anything, mock.Anything).
		Return(domain.Transfer{}, errors.New(" error"))
	repository.
		On("BeginTransaction",
			mock.Anything).
		Return(transaction, nil)

	accountRepository := new(domainMock.AccountRepository)
	accountRepository.
		On("Update",
			mock.Anything, mock.Anything, mock.Anything).
		Return(nil).Twice()
	accountRepository.
		On("GetByID",
			mock.Anything, mock.Anything).
		Return(domain.Account{}, nil).Twice()

	transfer, err := NewTransferController(repository, accountRepository).
		Process(context.Background(), domain.Transfer{})

	assert.Equal(t, domain.Transfer{}, transfer)
	assert.NotNil(t, err)
}

func TestController_CreateTransferUpdateAccountFailure(t *testing.T) {

	transaction := new(mocks.Transaction)
	transaction.
		On("Commit").
		Return(nil)
	transaction.
		On("Rollback").
		Return(nil)

	repository := new(domainMock.TransferRepository)
	repository.
		On("Process",
			mock.Anything, mock.Anything).
		Return(domain.Transfer{}, nil)
	repository.
		On("BeginTransaction",
			mock.Anything).
		Return(transaction, nil)

	accountRepository := new(domainMock.AccountRepository)
	accountRepository.
		On("Update",
			mock.Anything, mock.Anything, mock.Anything).
		Return(nil).Once()
	accountRepository.
		On("Update",
			mock.Anything, mock.Anything, mock.Anything).
		Return(errors.New("failed")).Once()
	accountRepository.
		On("GetByID",
			mock.Anything, mock.Anything).
		Return(domain.Account{}, nil).Twice()

	transfer, err := NewTransferController(repository, accountRepository).
		Process(context.Background(), domain.Transfer{})

	assert.Equal(t, domain.Transfer{}, transfer)
	assert.NotNil(t, err)
}

func TestController_CreateTransferGetAccountFailure(t *testing.T) {

	transaction := new(mocks.Transaction)
	transaction.
		On("Commit").
		Return(nil)

	repository := new(domainMock.TransferRepository)
	repository.
		On("Process",
			mock.Anything, mock.Anything).
		Return(domain.Transfer{}, nil)
	repository.
		On("BeginTransaction",
			mock.Anything).
		Return(transaction, nil)

	accountRepository := new(domainMock.AccountRepository)
	accountRepository.
		On("GetByID",
			mock.Anything, mock.Anything).
		Return(domain.Account{}, nil).Once()
	accountRepository.
		On("GetByID",
			mock.Anything, mock.Anything).
		Return(domain.Account{}, errors.New("failed")).Once()

	transfer, err := NewTransferController(repository, accountRepository).
		Process(context.Background(), domain.Transfer{})

	assert.Equal(t, domain.Transfer{}, transfer)
	assert.NotNil(t, err)
}

func TestController_CreateTransferGetAccountCombination2Failure(t *testing.T) {
	transaction := new(mocks.Transaction)
	transaction.
		On("Commit").
		Return(nil)

	repository := new(domainMock.TransferRepository)
	repository.
		On("Process",
			mock.Anything, mock.Anything).
		Return(domain.Transfer{}, nil)
	repository.
		On("BeginTransaction",
			mock.Anything).
		Return(transaction, nil)

	accountRepository := new(domainMock.AccountRepository)
	accountRepository.
		On("GetByID",
			mock.Anything, mock.Anything).
		Return(domain.Account{}, errors.New("failed")).Once()
	accountRepository.
		On("GetByID",
			mock.Anything, mock.Anything).
		Return(domain.Account{}, nil).Once()

	transfer, err := NewTransferController(repository, accountRepository).
		Process(context.Background(), domain.Transfer{})

	assert.Equal(t, domain.Transfer{}, transfer)
	assert.NotNil(t, err)
}

func TestController_GetTransfersSuccessfully(t *testing.T) {
	repository := new(domainMock.TransferRepository)
	repository.
		On("GetAll",
			mock.Anything).
		Return([]domain.Transfer{}, nil)

	accountRepository := new(domainMock.AccountRepository)

	createdAccount, err := NewTransferController(repository, accountRepository).GetAll(context.Background())

	assert.Equal(t, []domain.Transfer{}, createdAccount)
	assert.Equal(t, nil, err)
}

func TestController_GetTransfersFailure(t *testing.T) {
	repository := new(domainMock.TransferRepository)
	repository.
		On("GetAll",
			mock.Anything).
		Return([]domain.Transfer{}, errors.New(" failure to retrieve accounts"))

	accountRepository := new(domainMock.AccountRepository)

	createdAccount, err := NewTransferController(repository, accountRepository).GetAll(context.Background())

	assert.Equal(t, []domain.Transfer{}, createdAccount)
	assert.NotNil(t, err)
}
