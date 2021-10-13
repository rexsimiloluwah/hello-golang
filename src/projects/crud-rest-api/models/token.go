package models

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(user *User) (string, error) {
	var err error
	// JWT Claims
	jwtClaims := jwt.MapClaims{}
	jwtClaims["user_id"] = user.ID
	jwtClaims["exp"] = time.Now().Add(time.Minute * 45).Unix()

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)
	JWT_SECRET := os.Getenv("JWT_SECRET_KEY")
	token, err := jwtToken.SignedString(
		[]byte(JWT_SECRET),
	)
	if err != nil {
		return "", err
	}
	return token, nil
}

func VerifyToken(tokenString string) (jwt.Claims, error) {
	JWT_SECRET := os.Getenv("JWT_SECRET_KEY")
	tokenResponse, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWT_SECRET), nil
	})

	if err != nil {
		return nil, err
	}

	claims := tokenResponse.Claims
	if !tokenResponse.Valid {
		return nil, errors.New("Invalid or Expired token")
	}

	return claims, nil
}
