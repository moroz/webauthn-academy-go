package service

import (
	"github.com/alexedwards/argon2id"
	"github.com/gookit/validate"
	"github.com/jmoiron/sqlx"
	"github.com/moroz/webauthn-academy-go/store"
	"github.com/moroz/webauthn-academy-go/types"
)

type UserService struct {
	store store.UserStore
}

func NewUserService(db *sqlx.DB) UserService {
	return UserService{store.NewUserStore(db)}
}

func (s *UserService) RegisterUser(params types.NewUserParams) (*types.User, error, validate.Errors) {
	v := validate.Struct(params)

	if !v.Validate() {
		return nil, nil, v.Errors
	}

	passwordHash, err := argon2id.CreateHash(params.Password, argon2id.DefaultParams)
	if err != nil {
		return nil, err, nil
	}

	user, err := s.store.InsertUser(&types.User{
		Email:        params.Email,
		PasswordHash: passwordHash,
		DisplayName:  params.DisplayName,
	})

	if err != nil {
		return nil, err, nil
	}
	return user, nil, nil
}
