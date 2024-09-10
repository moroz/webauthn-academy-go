package service

import (
	"context"
	"errors"

	"github.com/alexedwards/argon2id"
	"github.com/gookit/validate"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/moroz/webauthn-academy-go/db/queries"
	"github.com/moroz/webauthn-academy-go/types"
)

func init() {
	validate.Config(func(opt *validate.GlobalOption) {
		opt.StopOnError = false
	})
}

type UserService struct {
	queries *queries.Queries
}

func NewUserService(db queries.DBTX) UserService {
	return UserService{queries.New(db)}
}

func (s *UserService) RegisterUser(ctx context.Context, params types.NewUserParams) (*queries.User, error, validate.Errors) {
	v := validate.Struct(params)

	if !v.Validate() {
		return nil, nil, v.Errors
	}

	passwordHash, err := argon2id.CreateHash(params.Password, argon2id.DefaultParams)
	if err != nil {
		return nil, err, nil
	}

	user, err := s.queries.InsertUser(ctx, queries.InsertUserParams{
		Email:        params.Email,
		PasswordHash: passwordHash,
		DisplayName:  params.DisplayName,
	})

	if err == nil {
		return user, nil, nil
	}

	// https://www.postgresql.org/docs/current/errcodes-appendix.html
	// Error 23505 `unique_violation` means that a unique constraint has prevented us from inserting
	// a duplicate value. Instead of returning a raw error, we return a handcrafted validation error
	// that we can later display in a form.
	if err, ok := err.(*pgconn.PgError); ok && err.Code == "23505" && err.ConstraintName == "users_email_key" {
		validationErrors := validate.Errors{}
		validationErrors.Add("Email", "unique", "has already been taken")
		return nil, nil, validationErrors
	}

	return nil, err, nil
}

var (
	UserNotFound    = errors.New("User not found")
	InvalidPassword = errors.New("Invalid password")
)

func (s *UserService) AuthenticateUserByEmailPassword(ctx context.Context, email, password string) (*queries.User, error) {
	user, err := s.queries.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if types.CheckUserPassword(user, password) {
		return user, nil
	}

	return nil, InvalidPassword
}
