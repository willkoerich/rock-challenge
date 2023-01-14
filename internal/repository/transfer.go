package repository

import (
	"context"
	"errors"
	"github.com/willkoerich/rock-challenge/internal/domain"
	"github.com/willkoerich/rock-challenge/internal/plataform/database"
	"time"
)

var (
	ErrTransferNotExist = errors.New("transfer row does not exist")
)

type TransferRepositoryDefault struct {
	driver database.Driver
}

func NewTransferRepository(driver database.Driver) domain.TransferRepository {
	return TransferRepositoryDefault{
		driver: driver,
	}
}

func (repository TransferRepositoryDefault) Process(ctx context.Context, transfer domain.Transfer) (domain.Transfer, error) {
	id, err := repository.driver.ExecuteInsertCommand(
		ctx,
		"INSERT INTO challenge.transfer(account_origin_id, account_destination_id, amount, created_at) values($1, $2, $3, $4) RETURNING id",
		transfer.AccountOriginID, transfer.AccountDestinationID, transfer.Amount, time.Now())
	if err != nil {
		return domain.Transfer{}, err
	}
	transfer.ID = id

	return transfer, nil
}

func (repository TransferRepositoryDefault) GetByID(ctx context.Context, id int) (domain.Transfer, error) {
	row := repository.driver.ExecuteQuerySingleElementCommand(ctx, "SELECT * FROM challenge.transfer where id = $1", id)
	var transfer domain.Transfer
	if row != nil {
		if err := row.Scan(&transfer.ID, &transfer.AccountOriginID, &transfer.AccountDestinationID, &transfer.Amount, &transfer.CreatedAt); err != nil {
			return domain.Transfer{}, err
		}
	} else {
		return domain.Transfer{}, ErrTransferNotExist
	}

	return transfer, nil
}

func (repository TransferRepositoryDefault) GetByAccountOriginID(ctx context.Context, accountOriginID int) ([]domain.Transfer, error) {
	rows, err := repository.driver.ExecuteQueryElementSetCommand(ctx, "SELECT * FROM challenge.transfer where accountOriginID = $1", accountOriginID)
	if err != nil {
		return []domain.Transfer{}, err
	}
	defer rows.Close()

	var all []domain.Transfer
	for rows.Next() {
		var transfer domain.Transfer
		if err := rows.Scan(&transfer.ID, &transfer.AccountOriginID, &transfer.AccountDestinationID, &transfer.Amount, &transfer.CreatedAt); err != nil {
			return []domain.Transfer{}, err
		}
		all = append(all, transfer)
	}
	return all, nil
}

func (repository TransferRepositoryDefault) GetAll(ctx context.Context) ([]domain.Transfer, error) {
	rows, err := repository.driver.ExecuteQueryElementSetCommand(ctx, "SELECT * FROM challenge.transfer ORDER BY id")
	if err != nil {
		return []domain.Transfer{}, err
	}
	defer rows.Close()

	var all []domain.Transfer
	for rows.Next() {
		var transfer domain.Transfer
		if err := rows.Scan(&transfer.ID, &transfer.AccountOriginID, &transfer.AccountDestinationID, &transfer.Amount, &transfer.CreatedAt); err != nil {
			return []domain.Transfer{}, err
		}
		all = append(all, transfer)
	}

	return all, nil
}
