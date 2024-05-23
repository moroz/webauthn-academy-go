package store

import (
	"github.com/jmoiron/sqlx"
	"github.com/moroz/webauthn-academy-go/types"
)

type UserTokenStore struct {
	db *sqlx.DB
}

func NewUserTokenStore(db *sqlx.DB) UserTokenStore {
	return UserTokenStore{db}
}

const insertUserTokenQuery = `insert into user_tokens (user_id, token, context) values ($1, $2, $3) returning *`

func (s *UserTokenStore) InsertToken(token *types.UserToken) (*types.UserToken, error) {
	var result types.UserToken
	err := s.db.Get(&result, insertUserTokenQuery, token.UserId, token.Token, token.Context)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

const getUserByTokenQuery = `select u.* from users u join user_tokens ut on u.id = ut.user_id where ut.token = $1 and ut.context = $2 and ut.inserted_at > (now() at time zone 'utc') - ($3 * interval '1 day')`

func (s *UserTokenStore) GetUserByToken(token []byte, context types.UserTokenContext, validityDays int) (*types.User, error) {
	var result types.User
	err := s.db.Get(&result, getUserByTokenQuery, token, context, validityDays)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
