package controller

import (
	"context"
	"errors"
	"fmt"
	"github.com/willkoerich/rock-challenge/internal/domain"
	"github.com/willkoerich/rock-challenge/internal/plataform/crypto"
)

type (
	LoginControllerDefault struct {
		Repository        domain.AccountRepository
		PasswordGenerator crypto.SecurePasswordService
	}
)

func NewLoginController(repository domain.AccountRepository, passwordGenerator crypto.SecurePasswordService) domain.LoginController {
	return LoginControllerDefault{
		Repository:        repository,
		PasswordGenerator: passwordGenerator,
	}
}

func (controller LoginControllerDefault) Authenticate(ctx context.Context, request domain.AuthenticationRequest) (domain.AuthenticationResponse, error) {
	account, err := controller.Repository.GetByCPF(ctx, request.CPF)
	if err != nil {
		return domain.AuthenticationResponse{}, fmt.Errorf("error getting account by CPF %s. Err: %s", request.CPF, err.Error())
	}

	err = controller.PasswordGenerator.Compare(account.Secret, request.Secret)
	if err != nil {
		return domain.AuthenticationResponse{}, errors.New("unauthorized")
	}

	return domain.AuthenticationResponse{
		Name:      account.Name,
		CPF:       account.CPF,
		AccountID: account.ID,
	}, nil
}
