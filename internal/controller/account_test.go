package controller

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/willkoerich/rock-challenge/internal/domain"
	domainMock "github.com/willkoerich/rock-challenge/internal/mocks/domain"
	cryptoMock "github.com/willkoerich/rock-challenge/internal/mocks/plataform/crypto"
	"testing"
)

func TestController_CreateAccountSuccessfully(t *testing.T) {
	repository := new(domainMock.AccountRepository)
	repository.
		On("Save",
			mock.Anything, mock.Anything).
		Return(domain.Account{ID: 1}, nil)

	generator := new(cryptoMock.SecurePasswordService)
	generator.
		On("Generate",
			mock.Anything).
		Return("encryptedPassword", nil)

	_, err := NewAccountController(repository, generator).Create(context.Background(), domain.CreateRequest{
		Name:    "Jon snow",
		CPF:     "11122233399",
		Secret:  "dasdasdas",
		Balance: 100,
	})

	assert.Nil(t, err)
}

func TestController_CreateAccountFailure(t *testing.T) {
	repository := new(domainMock.AccountRepository)
	repository.
		On("Save",
			mock.Anything, mock.Anything).
		Return(domain.Account{}, errors.New(" error to persist account"))

	generator := new(cryptoMock.SecurePasswordService)
	generator.
		On("Generate",
			mock.Anything).
		Return("encryptedPassword", nil)

	createdAccount, err := NewAccountController(repository, generator).Create(context.Background(), domain.CreateRequest{})

	assert.Equal(t, domain.Account{}, createdAccount)
	assert.NotNil(t, err)
}

func TestController_CreateAccountWhenPasswordGeneratorFailure(t *testing.T) {
	repository := new(domainMock.AccountRepository)
	repository.
		On("Save",
			mock.Anything, mock.Anything).
		Return(domain.Account{}, nil)

	generator := new(cryptoMock.SecurePasswordService)
	generator.
		On("Generate",
			mock.Anything).
		Return("", errors.New("failed"))

	recoveredAccount, err := NewAccountController(repository, generator).Create(context.Background(), domain.CreateRequest{})

	assert.Equal(t, domain.Account{}, recoveredAccount)
	assert.NotNil(t, err)
}

func TestController_GetAccountByIDSuccessfully(t *testing.T) {
	repository := new(domainMock.AccountRepository)
	repository.
		On("GetByID",
			mock.Anything, mock.Anything).
		Return(domain.Account{}, nil)

	generator := new(cryptoMock.SecurePasswordService)

	recoveredAccount, err := NewAccountController(repository, generator).GetByID(context.Background(), 1)

	assert.Equal(t, domain.Account{}, recoveredAccount)
	assert.Equal(t, nil, err)
}

func TestController_GetAccountByIDFailure(t *testing.T) {
	repository := new(domainMock.AccountRepository)
	repository.
		On("GetByID",
			mock.Anything, mock.Anything).
		Return(domain.Account{}, errors.New(" failure to retrieve account"))

	generator := new(cryptoMock.SecurePasswordService)

	recoveredAccount, err := NewAccountController(repository, generator).GetByID(context.Background(), 1)

	assert.Equal(t, domain.Account{}, recoveredAccount)
	assert.NotNil(t, err)
}

func TestController_GetAccountByCPFSuccessfully(t *testing.T) {
	repository := new(domainMock.AccountRepository)
	repository.
		On("GetByCPF",
			mock.Anything, mock.Anything).
		Return(domain.Account{}, nil)

	generator := new(cryptoMock.SecurePasswordService)

	recoveredAccount, err := NewAccountController(repository, generator).GetByCPF(context.Background(), "11122233344")

	assert.Equal(t, domain.Account{}, recoveredAccount)
	assert.Equal(t, nil, err)
}

func TestController_GetAccountByCPFFailure(t *testing.T) {
	repository := new(domainMock.AccountRepository)
	repository.
		On("GetByCPF",
			mock.Anything, mock.Anything).
		Return(domain.Account{}, errors.New(" failure to retrieve account"))

	generator := new(cryptoMock.SecurePasswordService)

	recoveredAccount, err := NewAccountController(repository, generator).GetByCPF(context.Background(), "11122233344")

	assert.Equal(t, domain.Account{}, recoveredAccount)
	assert.NotNil(t, err)
}

func TestController_GetAccountsSuccessfully(t *testing.T) {
	repository := new(domainMock.AccountRepository)
	repository.
		On("GetAll",
			mock.Anything).
		Return([]domain.Account{}, nil)

	generator := new(cryptoMock.SecurePasswordService)

	createdAccount, err := NewAccountController(repository, generator).GetAll(context.Background())

	assert.Equal(t, []domain.Account{}, createdAccount)
	assert.Equal(t, nil, err)
}

func TestController_GetAccountsFailure(t *testing.T) {
	repository := new(domainMock.AccountRepository)
	repository.
		On("GetAll",
			mock.Anything).
		Return([]domain.Account{}, errors.New(" failure to retrieve accounts"))

	generator := new(cryptoMock.SecurePasswordService)

	createdAccount, err := NewAccountController(repository, generator).GetAll(context.Background())

	assert.Equal(t, []domain.Account{}, createdAccount)
	assert.NotNil(t, err)
}
