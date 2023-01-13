package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/willkoerich/rock-challenge/internal/domain"
	"github.com/willkoerich/rock-challenge/internal/plataform/database"
	"strconv"

	"time"
)

var (
	ErrAccountNotExist = errors.New("account row does not exist")
)

type AccountRepositoryDefault struct {
	driver database.Driver
}

func NewAccountRepositoryDefault(driver database.Driver) domain.AccountRepository {
	return AccountRepositoryDefault{
		driver: driver,
	}
}

func (repository AccountRepositoryDefault) Save(ctx context.Context, account domain.Account) (domain.Account, error) {
	id, err := repository.driver.ExecuteInsertCommand(
		ctx,
		"INSERT INTO account(name, cpf, secret, balance, created_at) values($1, $2, $3, $4, $5) RETURNING id",
		account.Name, account.CPF, account.Secret, fmt.Sprintf("%f", account.Balance), time.Now())
	if err != nil {
		return domain.Account{}, err
	}
	account.ID = id

	return account, nil
}

func (repository AccountRepositoryDefault) GetByID(ctx context.Context, id int) (domain.Account, error) {
	row := repository.driver.ExecuteQuerySingleElementCommand(ctx, "SELECT * FROM account WHERE id = $1", strconv.Itoa(id))

	var account domain.Account
	if row != nil {
		if err := row.Scan(&account.ID, &account.Name, &account.CPF, &account.Secret, &account.Balance, &account.CreatedAt); err != nil {
			return domain.Account{}, err
		}
	} else {
		return domain.Account{}, ErrAccountNotExist
	}

	return account, nil
}

func (repository AccountRepositoryDefault) GetByCPF(ctx context.Context, cpf string) (domain.Account, error) {
	row := repository.driver.ExecuteQuerySingleElementCommand(
		ctx, "SELECT * FROM account WHERE cpf = $1", cpf)
	var account domain.Account
	if row != nil {
		if err := row.
			Scan(&account.ID, &account.Name, &account.CPF,
				&account.Secret, &account.Balance, &account.CreatedAt); err != nil {
			return domain.Account{}, err
		}
	} else {
		return domain.Account{}, ErrAccountNotExist
	}

	return account, nil
}

func (repository AccountRepositoryDefault) GetAll(ctx context.Context) ([]domain.Account, error) {
	rows, err := repository.driver.ExecuteQueryElementSetCommand(ctx, "SELECT * FROM account")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var all []domain.Account
	for rows.Next() {
		var account domain.Account
		if err := rows.Scan(&account.ID, &account.Name, &account.CPF,
			&account.Secret, &account.Balance, &account.CreatedAt); err != nil {
			return nil, err
		}
		all = append(all, account)
	}
	return all, nil
}
