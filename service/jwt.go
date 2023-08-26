package service

import (
	"errors"
	"fmt"
	"github.com/waretool/go-common/domain"
	"github.com/waretool/go-common/env"
	"github.com/waretool/go-common/logger"
	"time"

	"github.com/golang-jwt/jwt"
)

type jwtService struct {
	key []byte
}

func NewJwtService() domain.JwtService {
	key := env.GetEnv("JWT_SECRET_KEY", "")
	if key == "" {
		logger.Fatalf("env variable JWT_SECRET_KEY is missing but is mandatory")
	}
	return &jwtService{
		key: []byte(key),
	}
}

func (s *jwtService) Generate(consumer domain.Consumer) (string, error) {
	if consumer == nil {
		return "", errors.New("unable to generate a jwt from a nil consumer")
	}

	exp := env.GetEnv("JWT_EXP_TIME", 900)
	expirationTime := time.Now().Add(time.Duration(exp) * time.Second)

	claims := domain.Claims{
		Role:    consumer.GetRole(),
		Enabled: consumer.IsEnabled(),
		Local:   consumer.IsLocal(),
		StandardClaims: jwt.StandardClaims{
			Audience:  "",
			ExpiresAt: expirationTime.Unix(),
			Id:        "",
			IssuedAt:  time.Now().Unix(),
			Issuer:    env.GetEnv("APP_NAME", ""),
			Subject:   consumer.GetIdentifier(),
		},
	}

	logger.Debugf("generating jwt with claims %v", claims)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(s.key)
	if err != nil {
		return "", fmt.Errorf("unable to generate valid jwt due to: %s", err)
	}

	return tokenString, nil
}

func (s *jwtService) Valid(tokenString string) (*domain.Claims, bool) {
	token, err := jwt.ParseWithClaims(tokenString, &domain.Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return s.key, nil
	})

	if err != nil {
		return nil, false
	}

	if claims, ok := token.Claims.(*domain.Claims); ok && token.Valid {
		logger.Debugf("token %s is valid", token.Raw[:10]+"***")
		return claims, true
	} else {
		logger.Debugf("token %s... is not valid", token.Raw[:10]+"***")
		return nil, false
	}
}
