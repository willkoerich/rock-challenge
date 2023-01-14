package registry

import (
	"github.com/willkoerich/rock-challenge/internal/controller"
	"github.com/willkoerich/rock-challenge/internal/domain"
	"github.com/willkoerich/rock-challenge/internal/plataform/crypto"
	"github.com/willkoerich/rock-challenge/internal/plataform/database"
	"github.com/willkoerich/rock-challenge/internal/repository"
)

func NewAccountController(driver database.Driver, passwordGenerator crypto.SecurePasswordService) domain.AccountController {
	return controller.NewAccountController(NewAccountRepository(driver), passwordGenerator)
}

func NewAccountRepository(driver database.Driver) domain.AccountRepository {
	return repository.NewAccountRepository(driver)
}
