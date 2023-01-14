package domain

import (
	"context"
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
		Process(ctx context.Context, transfer Transfer) (Transfer, error)
		GetByID(ctx context.Context, id int) (Transfer, error)
		GetByAccountOriginID(ctx context.Context, accountOriginID int) ([]Transfer, error)
		GetAll(ctx context.Context) ([]Transfer, error)
	}

	TransferController interface {
		Process(context context.Context, transfer Transfer) (Transfer, error)
		GetAll(ctx context.Context) ([]Transfer, error)
	}
)
