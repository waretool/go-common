package domain

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type RoleSuite struct {
	suite.Suite
}

func TestUserSuite(t *testing.T) {
	suite.Run(t, new(RoleSuite))
}

func (suite *RoleSuite) TestRoleConst() {
	suite.Equal(Role(0), Admin)
}
