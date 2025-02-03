package auth

import (
	"errors"
	"math/rand"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Claims struct {
	UserID uuid.UUID `json:"userID"`
	jwt.RegisteredClaims
}

func (c *Claims) Valid() error {
	if c.ExpiresAt != nil && c.ExpiresAt.Time.Before(time.Now()) {
		return errors.New("token has expired")
	}
	return nil
}

func (auth *Auth) generateRandomSixDigit() int64 {
	otp := rand.Intn(999999)
	return int64(otp)
}

func (auth *Auth) GenerateTokens(userID uuid.UUID) (string, string, error) {
	accessTokenExpiry := time.Now().Add(auth.jwtExpiry)
	refreshTokenExpiry := time.Now().Add(auth.refreshExpiry)

	accessClaims := Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(accessTokenExpiry),
		},
	}

	refreshClaims := Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(refreshTokenExpiry),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)

	accessTokenString, err := accessToken.SignedString([]byte(auth.jwtSecret))
	if err != nil {
		return "", "", err
	}

	refreshTokenString, err := refreshToken.SignedString([]byte(auth.jwtSecret))
	if err != nil {
		return "", "", err
	}

	return accessTokenString, refreshTokenString, nil
}
