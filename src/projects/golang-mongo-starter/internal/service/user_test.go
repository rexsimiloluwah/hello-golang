package service

import (
	"testing"

	"github.com/rexsimiloluwah/hello-golang/src/projects/golang-mongo-starter/internal/config"
	usermock "github.com/rexsimiloluwah/hello-golang/src/projects/golang-mongo-starter/internal/tests/user"
	"github.com/rexsimiloluwah/hello-golang/src/projects/golang-mongo-starter/pkg/domain"
	"github.com/stretchr/testify/assert"
)

func TestCreateAccount_UserExists(t *testing.T) {
	cfg := &config.Settings{}
	userRepositoryMock := &usermock.UserRepositoryMock{}
	userSvc := NewUserService(cfg, userRepositoryMock)

	userRepositoryMock.EmailExistsMock = func(string) (bool, error) {
		return true, nil
	}

	// Mock a new user
	newUser := &domain.User{Email: "rexsimiloluwa@gmail.com", Username: "theblackdove", Password: "adetoyosi"}
	response := userSvc.CreateAccount(newUser)

	assert.Equal(t, "USER_EXISTS", response.Name)
}
