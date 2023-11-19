package crypto

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/suite"
)

type CryptoSuite struct {
	suite.Suite
	validKey string
}

func (suite *CryptoSuite) SetupTest() {
	suite.validKey = "1234567890abcdef"
}

func TestCryptoSuite(t *testing.T) {
	suite.Run(t, new(CryptoSuite))
}

func (suite *CryptoSuite) TestEncrypt() {
	text := "text to be encrypted"
	encText, _ := Encrypt(suite.validKey, text)
	match, _ := regexp.MatchString("^([A-Za-z0-9+/]{4})*([A-Za-z0-9+/]{3}=|[A-Za-z0-9+/]{2}==)?$", encText)
	suite.True(match)
}

func (suite *CryptoSuite) TestEncryptWithBadKey() {
	text := "text to be encrypted"
	_, err := Encrypt("key too short", text)
	suite.ErrorContains(err, "unable to encrypt data")
}

func (suite *CryptoSuite) TestDecrypt() {
	text := "text to be encrypted"
	encText, _ := Encrypt(suite.validKey, text)
	decText, _ := Decrypt(suite.validKey, encText)
	suite.Equal(text, decText)
}

func (suite *CryptoSuite) TestDecryptWithBadKey() {
	text := "text to be encrypted"
	encText, _ := Encrypt(suite.validKey, text)
	_, err := Decrypt("key too short", encText)
	suite.ErrorContains(err, "unable to decrypt data")
}

func (suite *CryptoSuite) TestDecryptTextNotBase64Encoded() {
	_, err := Decrypt(suite.validKey, "not valid base64 encoded text")
	suite.ErrorContains(err, "text is not base64 encoded")
}
