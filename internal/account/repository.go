package account

import (
	"context"
	"errors"
	"fmt"
	"github.com/willkoerich/rock-challenge/internal/plataform/database"
	"strconv"

	"time"
)

var (
	ErrAccountNotExist = errors.New("account row does not exist")
)

type RepositoryDefault struct {
	driver database.Driver
}

func NewRepositoryDefault(driver database.Driver) Repository {
	return RepositoryDefault{
		driver: driver,
	}
}

func (repository RepositoryDefault) Save(ctx context.Context, account Account) (Account, error) {
	id, err := repository.driver.ExecuteInsertCommand(
		ctx,
		"INSERT INTO account(name, cpf, secret, balance, created_at) values($1, $2, $3, $4, $5) RETURNING id",
		account.Name, account.CPF, account.Secret, fmt.Sprintf("%f", account.Balance), time.Now())
	if err != nil {
		return Account{}, err
	}
	account.ID = id

	return account, nil
}

func (repository RepositoryDefault) GetByID(ctx context.Context, id int) (Account, error) {
	row := repository.driver.ExecuteQuerySingleElementCommand(ctx, "SELECT * FROM account WHERE id = $1", strconv.Itoa(id))

	var account Account
	if row != nil {
		if err := row.Scan(&account.ID, &account.Name, &account.CPF, &account.Secret, &account.Balance, &account.CreatedAt); err != nil {
			return Account{}, err
		}
	} else {
		return Account{}, ErrAccountNotExist
	}

	return account, nil
}

func (repository RepositoryDefault) GetByCPF(ctx context.Context, cpf string) (Account, error) {
	row := repository.driver.ExecuteQuerySingleElementCommand(
		ctx, "SELECT * FROM account WHERE cpf = $1", cpf)
	var account Account
	if row != nil {
		if err := row.
			Scan(&account.ID, &account.Name, &account.CPF,
				&account.Secret, &account.Balance, &account.CreatedAt); err != nil {
			return Account{}, err
		}
	} else {
		return Account{}, ErrAccountNotExist
	}

	return account, nil
}

func (repository RepositoryDefault) GetAll(ctx context.Context) ([]Account, error) {
	rows, err := repository.driver.ExecuteQueryElementSetCommand(ctx, "SELECT * FROM account")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var all []Account
	for rows.Next() {
		var account Account
		if err := rows.Scan(&account.ID, &account.Name, &account.CPF,
			&account.Secret, &account.Balance, &account.CreatedAt); err != nil {
			return nil, err
		}
		all = append(all, account)
	}
	return all, nil
}
