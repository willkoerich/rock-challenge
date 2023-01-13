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

func TestController_AuthenticateSuccessfully(t *testing.T) {
	id := 1
	name := "Jon Doe"
	cpf := "11122233344"

	repository := new(domainMock.AccountRepository)
	repository.
		On("GetByCPF",
			mock.Anything, mock.Anything).
		Return(domain.Account{
			ID:     id,
			Name:   name,
			CPF:    cpf,
			Secret: "$2a$08$GLs4I6Soh5GZTNCrU.qTIO8igtz8dojRqr2CeOXxRVBIcJA3ZGS6i",
		}, nil)

	generator := new(cryptoMock.SecurePasswordService)
	generator.
		On("Compare",
			mock.Anything, mock.Anything).
		Return(nil)

	response, err := NewLoginController(repository, generator).
		Authenticate(context.Background(), domain.AuthenticationRequest{Secret: "will123"})

	assert.Equal(t, domain.AuthenticationResponse{
		Name:      name,
		CPF:       cpf,
		AccountID: id,
	}, response)
	assert.Nil(t, err)
}

func TestController_AuthenticateInvalidPassword(t *testing.T) {
	repository := new(domainMock.AccountRepository)
	repository.
		On("GetByCPF",
			mock.Anything, mock.Anything).
		Return(domain.Account{ID: 1, Secret: "$2a$08$GLs4I6Soh5GZTNCrU.qTIO8igtz8dojRqr2CeOXxRVBIcJA3ZGS6i"}, nil)

	generator := new(cryptoMock.SecurePasswordService)
	generator.
		On("Compare",
			mock.Anything, mock.Anything).
		Return(errors.New("invalid password"))

	response, err := NewLoginController(repository, generator).
		Authenticate(context.Background(), domain.AuthenticationRequest{Secret: "stone123"})

	assert.Equal(t, domain.AuthenticationResponse{}, response)
	assert.NotNil(t, err)
}

func TestController_AuthenticateFailure(t *testing.T) {
	repository := new(domainMock.AccountRepository)
	repository.
		On("GetByCPF",
			mock.Anything, mock.Anything).
		Return(domain.Account{}, errors.New(" error to persist account"))

	generator := new(cryptoMock.SecurePasswordService)
	generator.
		On("Compare",
			mock.Anything, mock.Anything).
		Return(nil)

	response, err := NewLoginController(repository, generator).Authenticate(context.Background(), domain.AuthenticationRequest{})

	assert.Equal(t, domain.AuthenticationResponse{}, response)
	assert.NotNil(t, err)
}
