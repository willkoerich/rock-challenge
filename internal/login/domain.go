package login

type (
	AuthenticationRequest struct {
		CPF    string `json:"cpf"`
		Secret string `json:"secret"`
	}
)
