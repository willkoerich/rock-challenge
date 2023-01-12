package account

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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

	result, err := NewRepositoryDefault(driver).
		Save(
			context.Background(),
			Account{
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

	result, err := NewRepositoryDefault(driver).
		Save(
			context.Background(),
			Account{
				Name:      "TestUser",
				CPF:       "00011122233",
				Secret:    "mySecret",
				Balance:   50000,
				CreatedAt: time.Now(),
			},
		)

	assert.Equal(t, Account{}, result)
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

	response, err := NewRepositoryDefault(driver).
		GetByID(
			context.Background(),
			-999,
		)

	assert.Equal(t, nil, err)
	assert.Equal(t, Account{}, response)
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

	response, err := NewRepositoryDefault(driver).
		GetByID(
			context.Background(),
			-999,
		)

	assert.Equal(t, ErrAccountNotExist, err)
	assert.Equal(t, Account{}, response)
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

	response, err := NewRepositoryDefault(driver).
		GetByID(
			context.Background(),
			-999,
		)

	assert.NotEqual(t, nil, err)
	assert.Equal(t, Account{}, response)
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

	response, err := NewRepositoryDefault(driver).GetAll(context.Background())

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

	response, err := NewRepositoryDefault(driver).GetAll(context.Background())

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

	response, err := NewRepositoryDefault(driver).GetAll(context.Background())

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

	response, err := NewRepositoryDefault(driver).GetAll(context.Background())

	assert.NotEqual(t, nil, err)
	assert.Equal(t, 0, len(response))
}
