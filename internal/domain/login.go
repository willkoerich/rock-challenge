package domain

import "context"

type (
	AuthenticationRequest struct {
		CPF    string `json:"cpf"`
		Secret string `json:"secret"`
	}

	AuthenticationResponse struct {
		Name      string `json:"name"`
		CPF       string `json:"cpf"`
		AccountID int    `json:"account_id"`
	}

	TokenGenerationResponse struct {
		AccessToken string `json:"access_token"`
	}

	LoginController interface {
		Authenticate(ctx context.Context, credential AuthenticationRequest) (AuthenticationResponse, error)
	}
)
