package domain

import (
	"context"
	"github.com/willkoerich/rock-challenge/internal/plataform/database"
	"time"
)

type (
	Transfer struct {
		ID                   int       `json:"id"`
		AccountOriginID      int       `json:"account_origin_id"`
		AccountDestinationID int       `json:"account_destination_id"`
		Amount               float64   `json:"amount"`
		CreatedAt            time.Time `json:"created_at"`
	}

	TransferRepository interface {
		Process(ctx context.Context, transaction database.Transaction, transfer Transfer) (Transfer, error)
		GetByID(ctx context.Context, id int) (Transfer, error)
		GetByAccountOriginID(ctx context.Context, accountOriginID int) ([]Transfer, error)
		GetAll(ctx context.Context) ([]Transfer, error)
		BeginTransaction(ctx context.Context) (database.Transaction, error)
	}

	TransferController interface {
		Process(context context.Context, transfer Transfer) (Transfer, error)
		GetAll(ctx context.Context) ([]Transfer, error)
	}
)
