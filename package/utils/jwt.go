package utils

import (
	"time"
	"typer/package/exceptions"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecretKey []byte

func init() {
	jwtSecretKey = []byte(GetEnv("AUTH_SECRET", ""))

	if len(jwtSecretKey) == 0 {
		panic("AUTH_SECRET environment variable is not set or is empty")
	}
}

type Claims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

var defaultTokenDuration = time.Hour * 24

func GenerateJWT(userID string, duration ...time.Duration) (string, error) {
	dur := defaultTokenDuration
	if len(duration) > 0 {
		dur = duration[0]
	}
	claims := Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(dur)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecretKey)
}

func ParseJWT(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, &exceptions.ClientError{Code: 401, Message: "Invalid signing method"}
		}
		return jwtSecretKey, nil
	})
	if err != nil {
		return "", &exceptions.ClientError{Code: 401, Message: "Invalid token"}
	}
	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return "", &exceptions.ClientError{Code: 401, Message: "Invalid token claims"}
	}
	return claims.UserID, nil
}
