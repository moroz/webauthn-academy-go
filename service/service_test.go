package service_test

import (
	"crypto/rand"
	"fmt"
	"os"
	"testing"

	"github.com/alexedwards/argon2id"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/moroz/webauthn-academy-go/store"
	"github.com/moroz/webauthn-academy-go/types"
	"github.com/stretchr/testify/suite"
)

type ServiceTestSuite struct {
	suite.Suite
	db *sqlx.DB
}

func (s *ServiceTestSuite) SetupTest() {
	conn := os.Getenv("TEST_DATABASE_URL")
	s.db = sqlx.MustConnect("postgres", conn)
	s.db.MustExec("truncate users cascade")
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
