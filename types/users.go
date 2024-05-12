package types

import "time"

type User struct {
	ID           int       `db:"id"`
	Email        string    `db:"email"`
	DisplayName  string    `db:"display_name"`
	PasswordHash string    `db:"password_hash"`
	InsertedAt   time.Time `db:"inserted_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}

type NewUserParams struct {
	Email                string `schema:"email" validate:"required|email"`
	DisplayName          string `schema:"displayName" validate:"required"`
	Password             string `schema:"password" validate:"required|min_len:8|max_len:80"`
	PasswordConfirmation string `schema:"passwordConfirmation" validate:"required|eq_field:Password"`
}
