package controller

import (
	"context"
	"fmt"
	"github.com/willkoerich/rock-challenge/internal/domain"
	"github.com/willkoerich/rock-challenge/internal/plataform/crypto"
	"time"
)

const (
	ErrorToEncryptSecretMessage   = "error encrypting account secret. Err: %s"
	ErrorToGetAccountByIDMessage  = "error getting account by id %v. Err: %s"
	ErrorToGetAccountByCPFMessage = "error getting account by cpf %v. Err: %s"
	ErrorToGetAccountsMessage     = "error getting accounts. Err: %s"
)

type (
	AccountControllerDefault struct {
		Repository        domain.AccountRepository
		PasswordGenerator crypto.SecurePasswordService
	}
)

func NewAccountController(repository domain.AccountRepository, passwordGenerator crypto.SecurePasswordService) domain.AccountController {
	return AccountControllerDefault{
		Repository:        repository,
		PasswordGenerator: passwordGenerator,
	}
}

func (controller AccountControllerDefault) Create(ctx context.Context, request domain.CreateRequest) (domain.Account, error) {

	encryptedPassword, err := controller.PasswordGenerator.Generate(request.Secret)
	if err != nil {
		return domain.Account{}, fmt.Errorf(ErrorToEncryptSecretMessage, err.Error())
	}
	account := domain.Account{
		Name:      request.Name,
		CPF:       request.CPF,
		Secret:    encryptedPassword,
		Balance:   request.Balance,
		CreatedAt: time.Now(),
	}
	return controller.Repository.Save(ctx, account)
}

func (controller AccountControllerDefault) GetByID(ctx context.Context, id int) (domain.Account, error) {
	account, err := controller.Repository.GetByID(ctx, id)
	if err != nil {
		err = fmt.Errorf(ErrorToGetAccountByIDMessage, id, err.Error())
	}
	return account, err
}

func (controller AccountControllerDefault) GetByCPF(ctx context.Context, cpf string) (domain.Account, error) {
	account, err := controller.Repository.GetByCPF(ctx, cpf)
	if err != nil {
		err = fmt.Errorf(ErrorToGetAccountByCPFMessage, cpf, err.Error())
	}
	return account, err
}

func (controller AccountControllerDefault) GetAll(ctx context.Context) ([]domain.Account, error) {
	account, err := controller.Repository.GetAll(ctx)
	if err != nil {
		err = fmt.Errorf(ErrorToGetAccountsMessage, err.Error())
	}
	return account, err
}
