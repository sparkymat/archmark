package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type Claims struct {
	jwt.StandardClaims
}

const (
	TokenExpiryLengthHours = 8
	Issuer                 = "archmark"
)

func ParseJWTString(jwtSecret string, tokenString string) (*string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(_ *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return nil, fmt.Errorf("unable to parse token. err: %w", err)
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, fmt.Errorf("unable to obtain claims. err: %w", err)
	}

	return &claims.Subject, nil
}

func GenerateJWTString(jwtSecret string, userID uuid.UUID) (string, error) {
	claims := Claims{
		StandardClaims: jwt.StandardClaims{
			Subject:   userID.String(),
			ExpiresAt: time.Now().Add(time.Hour * TokenExpiryLengthHours).Unix(),
			Issuer:    "archmark",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS384, claims)

	signedToken, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", fmt.Errorf("failed to sign token. err: %w", err)
	}

	return signedToken, nil
}
