package models

import "time"

type User struct {
	ID        int
	Email     string
	Password  string
	Type      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ResetPasswordObject struct {
	Email         string `json:"email"`
	OldPassword   string `json:"old_password"`
	NewPassword   string `json:"new_password"`
	ReNewPassword string `json:"renew_password"`
}
