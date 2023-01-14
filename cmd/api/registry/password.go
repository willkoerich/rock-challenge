package registry

import (
	"github.com/willkoerich/rock-challenge/internal/plataform/crypto"
)

func NewPasswordGenerator() crypto.SecurePasswordService {
	return crypto.NewBcryptSecurePasswordService()
}
