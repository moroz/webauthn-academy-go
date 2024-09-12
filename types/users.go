package types

import (
	"github.com/alexedwards/argon2id"
	"github.com/gookit/validate"
	"github.com/moroz/webauthn-academy-go/db/queries"
)

type NewUserParams struct {
	Email                string `schema:"email" validate:"required|email"`
	DisplayName          string `schema:"displayName" validate:"required"`
	Password             string `schema:"password" validate:"required|min_len:8|max_len:80"`
	PasswordConfirmation string `schema:"passwordConfirmation" validate:"required|eq_field:Password"`
}

func (p NewUserParams) Messages() map[string]string {
	return validate.MS{
		"required": "can't be blank",
		"email":    "is not a valid email address",
		"min_len":  "must be between 8 and 80 characters long",
		"max_len":  "must be between 8 and 80 characters long",
		"eq_field": "passwords do not match",
	}
}

func CheckUserPassword(user *queries.User, password string) bool {
	match, _ := argon2id.ComparePasswordAndHash(password, user.PasswordHash)
	return match
}
