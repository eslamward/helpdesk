package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {

	passwordByte, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(passwordByte), nil
}

func ComparePassword(hashedPassword, password string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
