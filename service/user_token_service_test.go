package service_test

import (
	"crypto/rand"
	"fmt"

	"github.com/alexedwards/argon2id"
	"github.com/jmoiron/sqlx"
	"github.com/moroz/webauthn-academy-go/service"
	"github.com/moroz/webauthn-academy-go/store"
	"github.com/moroz/webauthn-academy-go/types"
)

func generateEmail() string {
	var suffix [2]byte
	rand.Read(suffix[:])
	return fmt.Sprintf("user-%#x@example.com", suffix)
}

const PASSWORD = "foobar"

func insertUser(db *sqlx.DB) (*types.User, error) {
	hash, _ := argon2id.CreateHash(PASSWORD, argon2id.DefaultParams)

	params := &types.User{
		Email:        generateEmail(),
		PasswordHash: hash,
		DisplayName:  "Test User",
	}

	store := store.NewUserStore(db)
	return store.InsertUser(params)
}

func (s *ServiceTestSuite) TestGenerateUserSessionToken() {
	user, err := insertUser(s.db)
	s.NoError(err)
	srv := service.NewUserTokenService(s.db)
	token, err := srv.GenerateUserSessionToken(user)
	s.NoError(err)
	s.Len(token, 32)

	actualUser, err := srv.GetUserBySessionToken(token)
	s.NoError(err)
	s.Equal(user.ID, actualUser.ID)

	fakeToken := service.GenerateRandomToken()
	actualUser, err = srv.GetUserBySessionToken(fakeToken)
	s.Error(err)
	s.Nil(actualUser)
}
