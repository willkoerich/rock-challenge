package transfer

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	mocks "github.com/willkoerich/rock-challenge/internal/mocks/plataform/database"
	"testing"
	"time"
)

func TestCreateTransferSuccessfully(t *testing.T) {

	driver := new(mocks.Driver)
	driver.
		On("ExecuteInsertCommand",
			mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(-9999, nil)

	response, err := NewRepositoryDefault(driver).
		Save(
			context.Background(),
			Transfer{
				AccountOriginID:      -2222,
				AccountDestinationID: -1111,
				Amount:               2000,
				CreatedAt:            time.Now(),
			},
		)

	assert.Equal(t, -9999, response.ID)
	assert.Nil(t, err)
}

func TestCreateTransferWhenExecuteInsertCommandFails(t *testing.T) {

	driver := new(mocks.Driver)
	driver.
		On("ExecuteInsertCommand",
			mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(0, errors.New("error"))

	response, err := NewRepositoryDefault(driver).
		Save(
			context.Background(),
			Transfer{
				AccountOriginID:      -2222,
				AccountDestinationID: -1111,
				Amount:               2000,
				CreatedAt:            time.Now(),
			},
		)

	assert.Equal(t, Transfer{}, response)
	assert.NotNil(t, err)
}

func TestGetTransferByIDSuccessfully(t *testing.T) {

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
	assert.Equal(t, Transfer{}, response)
}

func TestGetTransferByIDWhenExecuteQuerySingleElementCommandFails(t *testing.T) {

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

	assert.Equal(t, ErrTransferNotExist, err)
	assert.Equal(t, Transfer{}, response)
}

func TestGetTransferByIDWhenResultScanFails(t *testing.T) {

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
	assert.Equal(t, Transfer{}, response)
}

func TestGetTransferByAccountOriginIDSuccessfully(t *testing.T) {

	results := new(mocks.Results)
	results.
		On("Scan", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
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
			mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(results, nil)

	response, err := NewRepositoryDefault(driver).
		GetByAccountOriginID(
			context.Background(),
			-999,
		)

	assert.Equal(t, nil, err)
	assert.Equal(t, []Transfer{{}}, response)
}

func TestGetTransferByAccountOriginIDWhenExecuteQuerySingleElementCommandFails(t *testing.T) {

	results := new(mocks.Results)
	results.
		On("Scan", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
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
			mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(nil, errors.New("failed"))

	response, err := NewRepositoryDefault(driver).
		GetByAccountOriginID(
			context.Background(),
			-999,
		)

	assert.Equal(t, errors.New("failed"), err)
	assert.Equal(t, []Transfer{}, response)
}

func TestGetTransferByAccountOriginIDWhenResultScanFails(t *testing.T) {

	results := new(mocks.Results)
	results.
		On("Scan", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(nil)
	results.
		On("Close").
		Return(errors.New("failed"))
	results.
		On("Next").
		Return(true).Once()
	results.
		On("Next").
		Return(false).Once()

	driver := new(mocks.Driver)
	driver.
		On("ExecuteQueryElementSetCommand",
			mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(results, nil)

	response, err := NewRepositoryDefault(driver).
		GetByAccountOriginID(
			context.Background(),
			-999,
		)

	assert.NotEqual(t, errors.New("failed"), err)
	assert.Equal(t, []Transfer{{}}, response)
}

func TestGetAllTransfersSuccessfully(t *testing.T) {

	results := new(mocks.Results)
	results.
		On("Scan", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
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

func TestGetAllTransfersWhenScanFails(t *testing.T) {

	results := new(mocks.Results)
	results.
		On("Scan", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
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

func TestGetAllTransfersWhenResultCloseFails(t *testing.T) {

	results := new(mocks.Results)
	results.
		On("Scan", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
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

func TestGetAllTransfersWhenExecuteQueryElementSetCommandFails(t *testing.T) {

	results := new(mocks.Results)
	results.
		On("Scan", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
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
		Return(nil, errors.New("error to retrieve transfers"))

	response, err := NewRepositoryDefault(driver).GetAll(context.Background())

	assert.NotEqual(t, nil, err)
	assert.Equal(t, 0, len(response))
}
