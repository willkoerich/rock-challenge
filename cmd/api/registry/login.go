package registry

import (
	"github.com/willkoerich/rock-challenge/internal/controller"
	"github.com/willkoerich/rock-challenge/internal/domain"
	"github.com/willkoerich/rock-challenge/internal/plataform/crypto"
	"github.com/willkoerich/rock-challenge/internal/plataform/database"
)

func NewLoginController(driver database.Driver, passwordGenerator crypto.SecurePasswordService) domain.LoginController {
	return controller.NewLoginController(NewAccountRepository(driver), passwordGenerator)
}
