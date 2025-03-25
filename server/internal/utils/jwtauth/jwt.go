package jwtauth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type CustomAccessClaims struct {
	UserID uint `json:"userId"`
	jwt.RegisteredClaims
}

func GenerateAccessToken(userID uint, accessSecret []byte) (string, error) {
	claims := CustomAccessClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(accessSecret)
}

type CustomRefreshClaims struct {
	UserID       uint `json:"userId"`
	TokenVersion uint `json:"tokenVersion"`
	jwt.RegisteredClaims
}

func GenerateRefreshToken(userID, tokenVersion uint, refreshSecret []byte) (string, error) {
	claims := CustomRefreshClaims{
		UserID:       userID,
		TokenVersion: tokenVersion,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(refreshSecret)
}
