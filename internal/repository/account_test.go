package repository

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/willkoerich/rock-challenge/internal/domain"
	mocks "github.com/willkoerich/rock-challenge/internal/mocks/plataform/database"
	"testing"
	"time"
)

func TestCreateAccountSuccessfully(t *testing.T) {

	driver := new(mocks.Driver)
	driver.
		On("ExecuteInsertCommand",
			mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(-9999, nil)

	result, err := NewAccountRepository(driver).
		Save(
			context.Background(),
			domain.Account{
				Name:      "TestUser",
				CPF:       "00011122233",
				Secret:    "mySecret",
				Balance:   50000,
				CreatedAt: time.Now(),
			},
		)

	assert.Equal(t, nil, err)
	assert.Equal(t, -9999, result.ID)
}

func TestCreateAccountWhenExecuteInsertCommandFails(t *testing.T) {

	driver := new(mocks.Driver)
	driver.
		On("ExecuteInsertCommand",
			mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(0, ErrAccountNotExist)

	result, err := NewAccountRepository(driver).
		Save(
			context.Background(),
			domain.Account{
				Name:      "TestUser",
				CPF:       "00011122233",
				Secret:    "mySecret",
				Balance:   50000,
				CreatedAt: time.Now(),
			},
		)

	assert.Equal(t, domain.Account{}, result)
	assert.Equal(t, ErrAccountNotExist, err)
}

func TestGetAccountSuccessfully(t *testing.T) {

	result := new(mocks.Result)
	result.
		On("Scan", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(nil)

	driver := new(mocks.Driver)
	driver.
		On("ExecuteQuerySingleElementCommand",
			mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(result)

	response, err := NewAccountRepository(driver).
		GetByID(
			context.Background(),
			-999,
		)

	assert.Equal(t, nil, err)
	assert.Equal(t, domain.Account{}, response)
}

func TestGetAccountWhenExecuteQuerySingleElementCommandFails(t *testing.T) {

	result := new(mocks.Result)
	result.
		On("Scan", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(nil)

	driver := new(mocks.Driver)
	driver.
		On("ExecuteQuerySingleElementCommand",
			mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(nil)

	response, err := NewAccountRepository(driver).
		GetByID(
			context.Background(),
			-999,
		)

	assert.Equal(t, ErrAccountNotExist, err)
	assert.Equal(t, domain.Account{}, response)
}

func TestGetAccountWhenResultScanFails(t *testing.T) {

	result := new(mocks.Result)
	result.
		On("Scan", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(errors.New("failure to Scan result"))

	driver := new(mocks.Driver)
	driver.
		On("ExecuteQuerySingleElementCommand",
			mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(result)

	response, err := NewAccountRepository(driver).
		GetByID(
			context.Background(),
			-999,
		)

	assert.NotEqual(t, nil, err)
	assert.Equal(t, domain.Account{}, response)
}

func TestGetByCPFAccountSuccessfully(t *testing.T) {

	result := new(mocks.Result)
	result.
		On("Scan", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(nil)

	driver := new(mocks.Driver)
	driver.
		On("ExecuteQuerySingleElementCommand",
			mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(result)

	response, err := NewAccountRepository(driver).
		GetByCPF(
			context.Background(),
			"11122233344",
		)

	assert.Equal(t, nil, err)
	assert.Equal(t, domain.Account{}, response)
}

func TestGetByCPFAccountWhenExecuteQuerySingleElementCommandFails(t *testing.T) {

	result := new(mocks.Result)
	result.
		On("Scan", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(nil)

	driver := new(mocks.Driver)
	driver.
		On("ExecuteQuerySingleElementCommand",
			mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(nil)

	response, err := NewAccountRepository(driver).
		GetByCPF(
			context.Background(),
			"11122233344",
		)

	assert.Equal(t, ErrAccountNotExist, err)
	assert.Equal(t, domain.Account{}, response)
}

func TestGetByCPFAccountWhenResultScanFails(t *testing.T) {

	result := new(mocks.Result)
	result.
		On("Scan", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(errors.New("failure to Scan result"))

	driver := new(mocks.Driver)
	driver.
		On("ExecuteQuerySingleElementCommand",
			mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(result)

	response, err := NewAccountRepository(driver).
		GetByCPF(
			context.Background(),
			"11122233344",
		)

	assert.NotEqual(t, nil, err)
	assert.Equal(t, domain.Account{}, response)
}

func TestGetAllAccountSuccessfully(t *testing.T) {

	results := new(mocks.Results)
	results.
		On("Scan", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(nil)
	results.
		On("Close").
		Return(nil)
	results.
		On("Next").
		Return(true).Once()
	results.
		On("Next").
		Return(false).Once()

	driver := new(mocks.Driver)
	driver.
		On("ExecuteQueryElementSetCommand",
			mock.Anything, mock.Anything, mock.Anything).
		Return(results, nil)

	response, err := NewAccountRepository(driver).GetAll(context.Background())

	assert.Equal(t, nil, err)
	assert.Equal(t, 1, len(response))
}

func TestGetAllAccountWhenScanFails(t *testing.T) {

	results := new(mocks.Results)
	results.
		On("Scan", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(errors.New("error to Scan result"))
	results.
		On("Close").
		Return(nil)
	results.
		On("Next").
		Return(true).Once()
	results.
		On("Next").
		Return(false).Once()

	driver := new(mocks.Driver)
	driver.
		On("ExecuteQueryElementSetCommand",
			mock.Anything, mock.Anything, mock.Anything).
		Return(results, nil)

	response, err := NewAccountRepository(driver).GetAll(context.Background())

	assert.NotEqual(t, nil, err)
	assert.Equal(t, 0, len(response))
}

func TestGetAllAccountWhenResultCloseFails(t *testing.T) {

	results := new(mocks.Results)
	results.
		On("Scan", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(nil)
	results.
		On("Close").
		Return(errors.New("error to Close result"))
	results.
		On("Next").
		Return(true).Once()
	results.
		On("Next").
		Return(false).Once()

	driver := new(mocks.Driver)
	driver.
		On("ExecuteQueryElementSetCommand",
			mock.Anything, mock.Anything, mock.Anything).
		Return(results, nil)

	response, err := NewAccountRepository(driver).GetAll(context.Background())

	assert.Equal(t, nil, err)
	assert.Equal(t, 1, len(response))
}

func TestGetAllAccountWhenExecuteQueryElementSetCommandFails(t *testing.T) {

	results := new(mocks.Results)
	results.
		On("Scan", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(nil)
	results.
		On("Close").
		Return(nil)
	results.
		On("Next").
		Return(true).Once()
	results.
		On("Next").
		Return(false).Once()

	driver := new(mocks.Driver)
	driver.
		On("ExecuteQueryElementSetCommand",
			mock.Anything, mock.Anything, mock.Anything).
		Return(nil, errors.New("error to retrieve accounts"))

	response, err := NewAccountRepository(driver).GetAll(context.Background())

	assert.NotEqual(t, nil, err)
	assert.Equal(t, 0, len(response))
}

func TestUpdateAccountSuccessfully(t *testing.T) {

	results := new(mocks.ExecResult)
	results.
		On("LastInsertId", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(0, nil)
	results.
		On("RowsAffected").
		Return(0, nil)

	driver := new(mocks.Driver)
	driver.
		On("Exec",
			mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(results, nil)

	err := NewAccountRepository(driver).
		Update(
			context.Background(),
			domain.Account{
				Name:      "TestUser",
				CPF:       "00011122233",
				Secret:    "mySecret",
				Balance:   50000,
				CreatedAt: time.Now(),
			},
		)

	assert.Equal(t, nil, err)
}

func TestUpdateAccountWhenExecCommandFails(t *testing.T) {

	results := new(mocks.ExecResult)
	results.
		On("LastInsertId", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(0, nil)
	results.
		On("RowsAffected").
		Return(0, nil)

	driver := new(mocks.Driver)
	driver.
		On("Exec",
			mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(results, ErrAccountNotExist)

	err := NewAccountRepository(driver).
		Update(
			context.Background(),
			domain.Account{
				Name:      "TestUser",
				CPF:       "00011122233",
				Secret:    "mySecret",
				Balance:   50000,
				CreatedAt: time.Now(),
			},
		)

	assert.Equal(t, ErrAccountNotExist, err)
}
