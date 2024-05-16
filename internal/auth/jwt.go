package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type Claims struct {
	jwt.StandardClaims
	UserID string `json:"username"`
}

const TokenExpiryLengthHours = 8
const Issuer = "oxbook"

func ParseJWTString(jwtSecret string, tokenString string) (*string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return nil, fmt.Errorf("unable to parse token. err: %w", err)
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, fmt.Errorf("unable to obtain claims. err: %w", err)
	}

	return &claims.UserID, nil
}

func GenerateJWTString(jwtSecret string, userID uuid.UUID) (string, error) {
	claims := Claims{
		UserID: userID.String(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * TokenExpiryLengthHours).Unix(),
			Issuer:    "oxbook",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS384, claims)

	signedToken, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", fmt.Errorf("failed to sign token. err: %w", err)
	}

	return signedToken, nil
}
