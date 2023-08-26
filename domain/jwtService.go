package domain

type JwtService interface {
	Generate(consumer Consumer) (string, error)
	Valid(tokenString string) (*Claims, bool)
}
