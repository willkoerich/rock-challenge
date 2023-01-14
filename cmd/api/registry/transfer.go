package registry

import (
	"github.com/willkoerich/rock-challenge/internal/controller"
	"github.com/willkoerich/rock-challenge/internal/domain"
	"github.com/willkoerich/rock-challenge/internal/plataform/database"
	"github.com/willkoerich/rock-challenge/internal/repository"
)

func NewTransferController(driver database.Driver, accountRepository domain.AccountRepository) domain.TransferController {
	return controller.NewTransferController(NewTransferRepository(driver), accountRepository)
}

func NewTransferRepository(driver database.Driver) domain.TransferRepository {
	return repository.NewTransferRepository(driver)
}
