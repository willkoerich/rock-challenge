package controller

import (
	"context"
	"fmt"
	"github.com/willkoerich/rock-challenge/internal/domain"
)

type (
	TransferControllerDefault struct {
		Repository domain.TransferRepository
	}
)

func NewTransferController(repository domain.TransferRepository) domain.TransferController {
	return TransferControllerDefault{repository}
}

func (controller TransferControllerDefault) Save(ctx context.Context, transfer domain.Transfer) (domain.Transfer, error) {
	transfer, err := controller.Repository.Save(ctx, transfer)
	if err != nil {
		err = fmt.Errorf("error creating transfer. Err: %s", err.Error())
	}
	return transfer, err
}

func (controller TransferControllerDefault) GetAll(ctx context.Context) ([]domain.Transfer, error) {
	transfer, err := controller.Repository.GetAll(ctx)
	if err != nil {
		err = fmt.Errorf("error getting accounts. Err: %s", err.Error())
	}
	return transfer, err
}
