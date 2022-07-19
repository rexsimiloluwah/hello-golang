package api

import (
	"fmt"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/rexsimiloluwah/hello-golang/src/projects/golang-mongo-starter/internal/config"
	"github.com/rexsimiloluwah/hello-golang/src/projects/golang-mongo-starter/internal/service"
)

type Middleware struct {
	cfg *config.Settings
}

type IMiddleware interface {
	Auth(next echo.HandlerFunc) echo.HandlerFunc
}

func NewMiddleware(cfg *config.Settings) IMiddleware {
	return &Middleware{
		cfg: cfg,
	}
}

func (m Middleware) Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			return echo.ErrUnauthorized
		}

		token, err := jwt.ParseWithClaims(tokenString, &service.JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(m.cfg.JwtSecretKey), nil
		})

		if err != nil {
			fmt.Println(err)
			return echo.ErrUnauthorized
		}

		claims, ok := token.Claims.(*service.JwtCustomClaims)

		if ok && token.Valid {
			// This should be email (rectify this in the JWT service file)
			c.Set("user", claims.Name)
			return next(c)
		} else {
			return echo.ErrUnauthorized
		}
	}
}
