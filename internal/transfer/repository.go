package transfer

import (
	"context"
	"errors"
	"fmt"
	"github.com/willkoerich/rock-challenge/internal/plataform/database"
	"strconv"
	"time"
)

var (
	ErrTransferNotExist = errors.New("transfer row does not exist")
)

type RepositoryDefault struct {
	driver database.Driver
}

func NewRepositoryDefault(driver database.Driver) Repository {
	return RepositoryDefault{
		driver: driver,
	}
}

func (repository RepositoryDefault) Save(ctx context.Context, transfer Transfer) (Transfer, error) {
	id, err := repository.driver.ExecuteInsertCommand(
		ctx,
		"INSERT INTO transfer(account_origin_id, account_destination_id, created_at) values($1, $2, $3, $4) RETURNING id",
		transfer.AccountOriginID, strconv.Itoa(transfer.AccountDestinationID), fmt.Sprintf("%f", transfer.Amount), time.Now())
	if err != nil {
		return Transfer{}, err
	}
	transfer.ID = id

	return transfer, nil
}

func (repository RepositoryDefault) GetByID(ctx context.Context, id int) (Transfer, error) {
	row := repository.driver.ExecuteQuerySingleElementCommand(ctx, "SELECT * FROM transfer where id = $1", id)
	var transfer Transfer
	if row != nil {
		if err := row.Scan(&transfer.ID, &transfer.AccountOriginID, &transfer.AccountDestinationID, &transfer.Amount, &transfer.CreatedAt); err != nil {
			return Transfer{}, err
		}
	} else {
		return Transfer{}, ErrTransferNotExist
	}

	return transfer, nil
}

func (repository RepositoryDefault) GetByAccountOriginID(ctx context.Context, accountOriginID int) ([]Transfer, error) {
	rows, err := repository.driver.ExecuteQueryElementSetCommand(ctx, "SELECT * FROM transfer where accountOriginID = $1", accountOriginID)
	if err != nil {
		return []Transfer{}, err
	}
	defer rows.Close()

	var all []Transfer
	for rows.Next() {
		var transfer Transfer
		if err := rows.Scan(&transfer.ID, &transfer.AccountOriginID, &transfer.AccountDestinationID, &transfer.Amount, &transfer.CreatedAt); err != nil {
			return []Transfer{}, err
		}
		all = append(all, transfer)
	}
	return all, nil
}

func (repository RepositoryDefault) GetAll(ctx context.Context) ([]Transfer, error) {
	rows, err := repository.driver.ExecuteQueryElementSetCommand(ctx, "SELECT * FROM transfer")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var all []Transfer
	for rows.Next() {
		var transfer Transfer
		if err := rows.Scan(&transfer.ID, &transfer.AccountOriginID, &transfer.AccountDestinationID, &transfer.Amount, &transfer.CreatedAt); err != nil {
			return nil, err
		}
		all = append(all, transfer)
	}
	return all, nil
}
