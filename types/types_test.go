package types_test

import (
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/suite"
)

type TypesTestSuite struct {
	suite.Suite
}

func TestStoreTestSuite(t *testing.T) {
	suite.Run(t, new(TypesTestSuite))
}
