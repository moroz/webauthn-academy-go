package service

import (
	"crypto/rand"

	"github.com/jmoiron/sqlx"
	"github.com/moroz/webauthn-academy-go/store"
	"github.com/moroz/webauthn-academy-go/types"
)

type UserTokenService struct {
	store store.UserTokenStore
}

func NewUserTokenService(db *sqlx.DB) UserTokenService {
	return UserTokenService{store.NewUserTokenStore(db)}
}

const TOKEN_RAND_SIZE = 32
const SESSION_VALIDITY_IN_DAYS = 60

func generateRandomToken() []byte {
	var buffer = make([]byte, TOKEN_RAND_SIZE)
	rand.Read(buffer)
	return buffer
}

func (s *UserTokenService) GenerateUserSessionToken(user *types.User) ([]byte, error) {
	token := generateRandomToken()

	userToken := &types.UserToken{
		UserId:  user.ID,
		Token:   token,
		Context: types.UserTokenContext_Session,
	}

	if _, err := s.store.InsertToken(userToken); err != nil {
		return nil, err
	}

	return token, nil
}

func (s *UserTokenService) GetUserBySessionToken(token []byte) (*types.User, error) {
	return s.store.GetUserByToken(token, types.UserTokenContext_Session, SESSION_VALIDITY_IN_DAYS)
}
