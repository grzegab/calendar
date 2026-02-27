package auth

import (
	"context"
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID string
}

type TokenVerifier interface {
	Verify(ctx context.Context, token string) (Claims, error)
}

type Verifier struct {
	keyFunc jwt.Keyfunc
}

func NewVerifier(keyFunc jwt.Keyfunc) *Verifier {
	return &Verifier{keyFunc: keyFunc}
}

func (v *Verifier) Verify(ctx context.Context, tokenStr string) (Claims, error) {
	token, err := jwt.Parse(tokenStr, v.keyFunc)
	if err != nil || !token.Valid {
		return Claims{}, errors.New("invalid token")
	}

	mapClaims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return Claims{}, errors.New("invalid token")
	}

	userID, _ := mapClaims["sub"].(string)

	return Claims{
		UserID: userID,
	}, nil
}

func HMACKeyFunc(secret []byte) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		// opcjonalnie: sprawdź alg
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}

		return secret, nil
	}
}
