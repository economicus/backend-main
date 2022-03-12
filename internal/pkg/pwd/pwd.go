package pwd

import (
	"golang.org/x/crypto/bcrypt"
	e "main/internal/core/error"
)

const cost = 12

func Hash(pwd []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(pwd, cost)
}

func Compare(nonHashed, hashed []byte) error {
	err := bcrypt.CompareHashAndPassword(hashed, nonHashed)
	if err != nil {
		return e.ErrWrongPassword
	}
	return nil
}
