package domain

import "github.com/golang-jwt/jwt"

type Claims struct {
	jwt.StandardClaims
	Role    Role `json:"role,omitempty"`
	Enabled bool `json:"enabled,omitempty"`
	Local   bool `json:"local,omitempty"`
}
