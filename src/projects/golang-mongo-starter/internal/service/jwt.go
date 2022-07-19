package service

import (
	"fmt"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/rexsimiloluwah/hello-golang/src/projects/golang-mongo-starter/internal/config"
)

type JwtServiceInterface interface {
	GenerateToken(name string, id string) (string, error)
	ValidateToken(tokenString string) (*jwt.Token, error)
}

// JWT claim extending the standard JWT claims
type JwtCustomClaims struct {
	Name string `json:"name"`
	ID   string `json:"id"`
	jwt.StandardClaims
}

type JwtService struct {
	cfg    *config.Settings
	issuer string
}

func NewJwtService(cfg *config.Settings) JwtServiceInterface {
	return &JwtService{
		cfg:    cfg,
		issuer: "rexsimiloluwa@gmail.com",
	}
}

func (jwtSrv *JwtService) GenerateToken(name string, id string) (string, error) {
	// initialize the JWT custom claims
	// obtain the Jwt expires time
	jwtExpiresIn, err := strconv.Atoi(jwtSrv.cfg.JwtExpiresIn)
	if err != nil {
		panic(err)
	}
	expiresAt := time.Now().Add(time.Hour * time.Duration(jwtExpiresIn)).Unix() // expires in 24h
	jwtClaims := &JwtCustomClaims{
		name,
		id,
		jwt.StandardClaims{
			ExpiresAt: expiresAt,
			Issuer:    jwtSrv.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)

	// Generate encoded token using the secret signing key
	t, err := token.SignedString([]byte(jwtSrv.cfg.JwtSecretKey))
	if err != nil {
		return "", err
	}
	return t, nil
}

func (jwtSrv *JwtService) ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Signing method validation
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			// Signing method validation
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// Return the secret signing key if otherwise
		return []byte(jwtSrv.cfg.JwtSecretKey), nil
	})
}
