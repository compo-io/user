package user

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestUser(t *testing.T) {
	suite.Run(t, new(UserTestSuite))
}

type UserTestSuite struct {
	suite.Suite
}


