package service

import (
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

func (s *UserService) RegisterUser(params types.NewUserParams) (*types.User, error) {

}
