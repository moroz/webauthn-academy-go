package service_test

import (
	"strings"

	"github.com/alexedwards/argon2id"
	"github.com/moroz/webauthn-academy-go/service"
	"github.com/moroz/webauthn-academy-go/store"
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

func (s *ServiceTestSuite) TestRegisterUserWithInvalidParams() {
	params := types.NewUserParams{
		Email:                "invalid",
		DisplayName:          "Example User",
		Password:             "short",
		PasswordConfirmation: "not matching",
	}

	srv := service.NewUserService(s.db)
	user, err, validationErrors := srv.RegisterUser(params)
	s.NoError(err)
	s.Nil(user)
	msg := validationErrors.FieldOne("Email")
	s.Equal("is not a valid email address", msg)
	msg = validationErrors.FieldOne("Password")
	s.Equal("must be between 8 and 80 characters long", msg)
	msg = validationErrors.FieldOne("PasswordConfirmation")
	s.Contains(msg, "do not match")
}

func (s *ServiceTestSuite) TestRegisterUserWithDuplicateEmail() {
	store := store.NewUserStore(s.db)
	user, err := store.InsertUser(&types.User{
		Email:        "duplicate@email.com",
		PasswordHash: "test",
		DisplayName:  "John Smith",
	})

	s.NoError(err)

	srv := service.NewUserService(s.db)

	params := types.NewUserParams{
		Email:                user.Email,
		DisplayName:          "Other User",
		Password:             "foobar123123",
		PasswordConfirmation: "foobar123123",
	}
	user, err, validationErrors := srv.RegisterUser(params)
	s.Nil(user)
	s.Nil(err)
	msg := validationErrors.FieldOne("Email")
	s.Equal("has already been taken", msg)
}

func (s *ServiceTestSuite) TestAuthenticateUserByEmailPassword() {
	user, err := insertUser(s.db)
	s.NoError(err)

	examples := []struct {
		email    string
		password string
		expected bool
	}{
		{user.Email, PASSWORD, true},
		{strings.ToUpper(user.Email), PASSWORD, true},
		{user.Email, "invalid", false},
		{strings.ToUpper(user.Email), "invalid", false},
		{"invalid@example.com", PASSWORD, false},
	}
	for _, example := range examples {
		srv := service.NewUserService(s.db)
		actual, err := srv.AuthenticateUserByEmailPassword(example.email, example.password)
		if example.expected {
			s.NoError(err)
			s.Equal(user.ID, actual.ID)
		} else {
			s.Error(err)
			s.Nil(actual)
		}
	}
}
