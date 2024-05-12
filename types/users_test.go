package types_test

import (
	"github.com/gookit/validate"
	"github.com/moroz/webauthn-academy-go/types"
)

func (s *TypesTestSuite) TestValidateNewUserValidParams() {
	params := types.NewUserParams{
		Email:                "user@example.com",
		DisplayName:          "Example User",
		Password:             "foobar123",
		PasswordConfirmation: "foobar123",
	}

	v := validate.Struct(params)

	s.True(v.Validate())
}

func (s *TypesTestSuite) TestValidateNewUserWithoutEmail() {
	params := types.NewUserParams{
		Email:                "",
		DisplayName:          "Example User",
		Password:             "foobar123",
		PasswordConfirmation: "foobar123",
	}

	v := validate.Struct(params)

	s.False(v.Validate())
	actual := v.Errors.Field("Email")["required"]
	s.Contains(actual, "is required")
}
