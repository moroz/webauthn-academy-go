package service_test

import (
	"github.com/alexedwards/argon2id"
	"github.com/moroz/webauthn-academy-go/service"
	"github.com/moroz/webauthn-academy-go/types"
)

func (s *ServiceTestSuite) TestRegisterUser() {
	params := types.NewUserParams{
		Email:                "registration@example.com",
		DisplayName:          "Example User",
		Password:             "foobar123123",
		PasswordConfirmation: "foobar123123",
	}

	srv := service.NewUserService(s.db)
	user, err, _ := srv.RegisterUser(params)
	s.NoError(err)
	s.Equal(params.Email, user.Email)
	s.Equal(params.DisplayName, user.DisplayName)

	match, err := argon2id.ComparePasswordAndHash(params.Password, user.PasswordHash)
	s.True(match)
}
