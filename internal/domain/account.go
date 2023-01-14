package domain

import (
	"context"
	"time"
)

type (
	CreateRequest struct {
		Name    string  `json:"name"`
		CPF     string  `json:"cpf"`
		Secret  string  `json:"secret"`
		Balance float64 `json:"balance"`
	}

	Account struct {
		ID        int       `json:"id"`
		Name      string    `json:"name"`
		CPF       string    `json:"cpf"`
		Secret    string    `json:"secret,omitempty"`
		Balance   float64   `json:"balance"`
		CreatedAt time.Time `json:"create_at"`
	}

	BalanceResponse struct {
		Balance float64 `json:"balance"`
	}

	AccountRepository interface {
		Save(ctx context.Context, account Account) (Account, error)
		GetByID(ctx context.Context, id int) (Account, error)
		GetByCPF(ctx context.Context, cpf string) (Account, error)
		GetAll(ctx context.Context) ([]Account, error)
		Update(ctx context.Context, account Account) error
	}

	AccountController interface {
		Create(ctx context.Context, account CreateRequest) (Account, error)
		GetByID(ctx context.Context, id int) (Account, error)
		GetByCPF(ctx context.Context, cpf string) (Account, error)
		GetAll(ctx context.Context) ([]Account, error)
	}
)
