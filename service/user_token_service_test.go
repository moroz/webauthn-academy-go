package service_test

import (
	"context"

	"github.com/moroz/webauthn-academy-go/service"
)

func (s *ServiceTestSuite) TestGenerateUserSessionToken() {
	user, err := insertUser(s.db)
	s.NoError(err)
	srv := service.NewUserTokenService(s.db)
	token, err := srv.GenerateUserSessionToken(context.Background(), user)
	s.NoError(err)
	s.Len(token, 32)

	actualUser, err := srv.GetUserBySessionToken(context.Background(), token)
	s.NoError(err)
	s.Equal(user.ID, actualUser.ID)

	fakeToken := service.GenerateRandomToken()
	actualUser, err = srv.GetUserBySessionToken(context.Background(), fakeToken)
	s.Error(err)
	s.Zero(actualUser.ID)
}
