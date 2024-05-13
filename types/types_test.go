package types_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type TypesTestSuite struct {
	suite.Suite
}

func TestStoreTestSuite(t *testing.T) {
	suite.Run(t, new(TypesTestSuite))
}
