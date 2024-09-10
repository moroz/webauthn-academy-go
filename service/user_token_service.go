package service

import (
	"context"
	"crypto/rand"

	"github.com/moroz/webauthn-academy-go/db/queries"
	"github.com/moroz/webauthn-academy-go/types"
)

type UserTokenService struct {
	queries *queries.Queries
}

func NewUserTokenService(db queries.DBTX) UserTokenService {
	return UserTokenService{queries.New(db)}
}

const TOKEN_RAND_SIZE = 32
const SESSION_VALIDITY_IN_DAYS = 60

func GenerateRandomToken() []byte {
	var buffer = make([]byte, TOKEN_RAND_SIZE)
	rand.Read(buffer)
	return buffer
}

func (s *UserTokenService) GenerateUserSessionToken(ctx context.Context, user *queries.User) ([]byte, error) {
	token := GenerateRandomToken()

	userToken := queries.InsertUserTokenParams{
		UserID:  user.ID,
		Token:   token,
		Context: types.UserTokenContext_Session,
	}

	if _, err := s.queries.InsertUserToken(ctx, userToken); err != nil {
		return nil, err
	}

	return token, nil
}

func (s *UserTokenService) GetUserBySessionToken(ctx context.Context, token []byte) (*queries.User, error) {
	return s.queries.GetUserByToken(ctx, queries.GetUserByTokenParams{
		Token:        token,
		Context:      types.UserTokenContext_Session,
		ValidityDays: SESSION_VALIDITY_IN_DAYS,
	})
}
