package auth

import (
	"context"
	"errors"
	"github/grzegab/calendar/internal/users/infrastructure/jwt_generator"

	"github.com/golang-jwt/jwt/v5"
)

type TokenVerifier interface {
	Verify(ctx context.Context, token string) (jwt_generator.Claims, error)
}

type Verifier struct {
	keyFunc jwt.Keyfunc
}

func NewVerifier(keyFunc jwt.Keyfunc) *Verifier {
	return &Verifier{keyFunc: keyFunc}
}

func (v *Verifier) Verify(ctx context.Context, tokenStr string) (jwt_generator.Claims, error) {
	token, err := jwt.Parse(tokenStr, v.keyFunc)
	if err != nil || !token.Valid {
		return jwt_generator.Claims{}, errors.New("invalid token")
	}

	mapClaims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return jwt_generator.Claims{}, errors.New("invalid token")
	}

	userID, _ := mapClaims["sub"].(string)

	return jwt_generator.Claims{
		UserID: userID,
	}, nil
}

func HMACKeyFunc(secret []byte) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}

		return secret, nil
	}
}
