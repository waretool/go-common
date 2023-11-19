package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/waretool/go-common/env"

	"github.com/sirupsen/logrus"
	"golang.org/x/exp/slices"
)

var CipherKey string

func init() {
	CipherKey = env.GetEnv("CIPHER_KEY", "super-secret-key")
	if !slices.Contains([]int{16, 32, 64}, len(CipherKey)) {
		panic("invalid CIPHER_KEY value. it must be a 16, 32 or 64 byte compatible string")
	}
}

func Encrypt(key, text string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		logrus.Errorf("encrypt error. unable to create aes block cipher caused by: %s", err.Error())
		return "", fmt.Errorf("unable to encrypt data")
	}
	plaintext := []byte(text)
	iv := generateIv()
	cfb := cipher.NewCFBEncrypter(block, iv)
	ciphertext := make([]byte, len(plaintext))
	cfb.XORKeyStream(ciphertext, plaintext)

	return encodeBase64(append(iv, ciphertext...)), nil
}

func Decrypt(key, text string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		logrus.Errorf("decrypt error. unable to create aes block cipher caused by: %s", err.Error())
		return "", fmt.Errorf("unable to decrypt data")
	}
	ciphertext, err := decodeBase64(text)
	if err != nil {
		logrus.Errorf("unable to decode text due to: %s", err.Error())
		return "", fmt.Errorf("text is not base64 encoded")
	}
	iv := ciphertext[:16]
	encryptedText := ciphertext[16:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	plaintext := make([]byte, len(encryptedText))
	cfb.XORKeyStream(plaintext, encryptedText)

	return string(plaintext), nil
}

func generateIv() []byte {
	bytes := make([]byte, 16)
	rand.Read(bytes)
	return bytes
}

func decodeBase64(s string) ([]byte, error) {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func encodeBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}
