package password

import (
	"golang.org/x/crypto/bcrypt"
	"movie_api/helper"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	helper.PrintError(err)

	return string(bytes), nil
}
