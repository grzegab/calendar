package jwt_generator

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	UserID string `json:"sub"`
	jwt.RegisteredClaims
}

type JwtGenerator struct {
	secret string
}

func NewJwtGenerator(secret string) JwtGenerator {
	return JwtGenerator{secret: secret}
}

func (j JwtGenerator) Generate(c Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	signedToken, err := token.SignedString([]byte(j.secret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
