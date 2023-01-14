package controller

import (
	"context"
	"errors"
	"fmt"
	"github.com/willkoerich/rock-challenge/internal/domain"
)

const (
	ErrorToRetrieveOriginAccountMessage      = "unable to retrieve of origin account. Err: %s"
	ErrorToRetrieveDestinationAccountMessage = "unable to retrieve of destination account. Err: %s"
	ErrorOriginWithoutFoundsMessage          = "unable to transfer required amount as origin account doesn't have available founds"
	ErrorToCreateTransferMessage             = "error creating transfer. Err: %s"
)

type (
	TransferControllerDefault struct {
		Repository        domain.TransferRepository
		AccountRepository domain.AccountRepository
	}
)

func NewTransferController(repository domain.TransferRepository, accountRepository domain.AccountRepository) domain.TransferController {
	return TransferControllerDefault{
		Repository:        repository,
		AccountRepository: accountRepository,
	}
}

func (controller TransferControllerDefault) Process(ctx context.Context, transfer domain.Transfer) (domain.Transfer, error) {

	origin, err := controller.AccountRepository.GetByID(ctx, transfer.AccountOriginID)
	if err != nil {
		return domain.Transfer{}, fmt.Errorf(ErrorToRetrieveOriginAccountMessage, err.Error())
	}
	destination, err := controller.AccountRepository.GetByID(ctx, transfer.AccountDestinationID)
	if err != nil {
		return domain.Transfer{}, fmt.Errorf(ErrorToRetrieveDestinationAccountMessage, err.Error())
	}

	if origin.Balance < transfer.Amount {
		return domain.Transfer{}, errors.New(ErrorOriginWithoutFoundsMessage)
	}

	origin.Balance -= transfer.Amount
	destination.Balance += transfer.Amount
	controller.AccountRepository.Update(ctx, origin)
	controller.AccountRepository.Update(ctx, origin)

	transfer, err = controller.Repository.Process(ctx, transfer)
	if err != nil {
		err = fmt.Errorf(ErrorToCreateTransferMessage, err.Error())
	}
	return transfer, err
}

func (controller TransferControllerDefault) GetAll(ctx context.Context) ([]domain.Transfer, error) {
	return controller.Repository.GetAll(ctx)
}
