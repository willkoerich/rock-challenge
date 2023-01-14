package crypto

import (
	"golang.org/x/crypto/bcrypt"
)

const COST = 8

type (
	SecurePasswordService interface {
		Generate(password string) (string, error)
		Compare(hashedPassword, password string) error
	}

	BcryptSecurePasswordService struct{}
)

func NewBcryptSecurePasswordService() SecurePasswordService {
	return BcryptSecurePasswordService{}
}

func (service BcryptSecurePasswordService) Generate(password string) (string, error) {
	result, err := bcrypt.GenerateFromPassword([]byte(password), COST)
	if err != nil {
		return "", err
	}
	return string(result), err
}

func (service BcryptSecurePasswordService) Compare(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
