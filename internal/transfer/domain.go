package transfer

import (
	"context"
	"time"
)

type (
	Transfer struct {
		ID                   int
		AccountOriginID      int
		AccountDestinationID int
		Amount               float64
		CreatedAt            time.Time
	}

	Repository interface {
		Save(ctx context.Context, transfer Transfer) (Transfer, error)
		GetByID(ctx context.Context, id int) ([]Transfer, error)
		GetByAccountOriginID(ctx context.Context, accountOriginID int) ([]Transfer, error)
		GetAll(ctx context.Context) ([]Transfer, error)
	}
)
