package auth

import (
	"errors"

	"github.com/eslamward/helpdesk/models"
)

func userValidation(user models.User) error {
	if user.Email == "" || user.Password == "" || user.Type == "" {
		return errors.New("one or more field is empty")
	}
	return nil
}
