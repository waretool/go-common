package domain

import (
	"encoding/json"
	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ClaimsSuite struct {
	suite.Suite
}

func TestClaimsSuite(t *testing.T) {
	suite.Run(t, new(ClaimsSuite))
}

func (suite *ClaimsSuite) TestClaimsTag() {
	claims := Claims{
		StandardClaims: jwt.StandardClaims{},
		Role:           1,
	}

	bytes, _ := json.Marshal(claims)
	expected := `{ "role":1 }`
	suite.JSONEq(expected, string(bytes))
}

func (suite *ClaimsSuite) TestClaimsEmptyTag() {
	claims := Claims{}

	bytes, _ := json.Marshal(claims)
	suite.JSONEq("{}", string(bytes))
}
