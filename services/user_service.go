package services

import (
	"github.com/alexedwards/argon2id"
	"github.com/gookit/validate"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/moroz/webauthn-academy-go/db/queries"

	"context"
)

type UserService struct {
	queries *queries.Queries
}

func NewUserService(db queries.DBTX) UserService {
	return UserService{queries: queries.New(db)}
}

type RegisterUserParams struct {
	Email                string `validate:"required|email"`
	DisplayName          string `validate:"required"`
	Password             string `validate:"required|min_len:8|max_len:64"`
	PasswordConfirmation string `validate:"eq_field:Password" message:"passwords do not match"`
}

func init() {
	validate.AddGlobalMessages(map[string]string{
		"required": "can't be blank",
		"email":    "is not a valid email address",
		"min_len":  "must be at least %d characters long",
		"max_len":  "must be at most %d characters long",
	})
}

func (us *UserService) RegisterUser(ctx context.Context, params RegisterUserParams) (*queries.User, error) {
	if v := validate.Struct(&params); !v.Validate() {
		return nil, v.Errors
	}

	hash, err := argon2id.CreateHash(params.Password, argon2id.DefaultParams)
	if err != nil {
		return nil, err
	}

	user, err := us.queries.InsertUser(ctx, queries.InsertUserParams{
		Email:        params.Email,
		DisplayName:  params.DisplayName,
		PasswordHash: hash,
	})

	// intercept "unique_violation" errors on the email column
	if err, ok := err.(*pgconn.PgError); ok && err.Code == "23505" && err.ConstraintName == "users_email_key" {
		return nil, validate.Errors{"Email": {"unique": "has already been taken"}}
	}

	return user, err
}
