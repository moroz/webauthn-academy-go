package service_test

import (
	"context"
	"crypto/rand"
	"fmt"
	"os"
	"testing"

	"github.com/alexedwards/argon2id"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	"github.com/moroz/webauthn-academy-go/db/queries"
	"github.com/stretchr/testify/suite"
)

type ServiceTestSuite struct {
	suite.Suite
	db queries.DBTX
}

func (s *ServiceTestSuite) SetupTest() {
	conn := os.Getenv("TEST_DATABASE_URL")
	db, err := pgxpool.New(context.Background(), conn)
	s.NoError(err)
	s.db = db
	_, err = s.db.Exec(context.Background(), "truncate users cascade")
	s.NoError(err)
}

func TestServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ServiceTestSuite))
}

func generateEmail() string {
	var suffix [2]byte
	rand.Read(suffix[:])
	return fmt.Sprintf("user-%#x@example.com", suffix)
}

const PASSWORD = "foobar"

func insertUser(db queries.DBTX) (*queries.User, error) {
	hash, _ := argon2id.CreateHash(PASSWORD, argon2id.DefaultParams)

	params := queries.InsertUserParams{
		Email:        generateEmail(),
		PasswordHash: hash,
		DisplayName:  "Test User",
	}

	return queries.New(db).InsertUser(context.Background(), params)
}
