package crypto

import (
	"golang.org/x/crypto/bcrypt"
)

type (
	SecurePasswordService interface {
		Generate(password []byte, cost int) ([]byte, error)
		Compare(hashedPassword, password []byte) error
	}

	BcryptSecurePasswordService struct{}
)

func (service BcryptSecurePasswordService) Generate(password []byte, cost int) ([]byte, error) {
	return bcrypt.GenerateFromPassword(password, cost)
}

func (service BcryptSecurePasswordService) Compare(hashedPassword, password []byte) error {
	return bcrypt.CompareHashAndPassword(hashedPassword, password)
}
