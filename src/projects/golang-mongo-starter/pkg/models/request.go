package models

import (
	"github.com/labstack/echo/v4"
	"github.com/rexsimiloluwah/hello-golang/src/projects/golang-mongo-starter/pkg/domain"
)

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func ValidateRegisterRequest(c echo.Context) (*domain.User, *Error) {
	registerRequest := new(RegisterRequest)
	if err := c.Bind(registerRequest); err != nil {
		return nil, BindError()
	}

	var validationErrors []string

	if len(registerRequest.Password) < 8 {
		validationErrors = append(validationErrors, "Password must be at least 8 characters.")
	}
	if len(registerRequest.Username) < 3 {
		validationErrors = append(validationErrors, "Username must be at least 3 characters.")
	}
	if len(validationErrors) > 0 {
		return nil, ValidationError(validationErrors)
	}

	return &domain.User{
		Username: registerRequest.Username,
		Password: registerRequest.Password,
		Email:    registerRequest.Email,
	}, nil
}

func ValidateLoginRequest(c echo.Context) (*domain.User, *Error) {
	loginRequest := new(LoginRequest)
	if err := c.Bind(loginRequest); err != nil {
		return nil, BindError()
	}

	var validationErrors []string

	if len(loginRequest.Password) < 8 {
		validationErrors = append(validationErrors, "Password must be at least 8 characters.")
	}

	if len(validationErrors) > 0 {
		return nil, ValidationError(validationErrors)
	}

	return &domain.User{
		Email:    loginRequest.Email,
		Password: loginRequest.Password,
	}, nil
}
