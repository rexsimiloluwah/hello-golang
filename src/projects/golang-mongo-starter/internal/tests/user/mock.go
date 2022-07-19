package user_test

import "github.com/rexsimiloluwah/hello-golang/src/projects/golang-mongo-starter/pkg/domain"

// Create a user repository mock
type UserRepositoryMock struct {
	CreateAccountMock   func(user *domain.User) error
	EmailExistsMock     func(email string) (bool, error)
	FindUserByEmailMock func(email string) (*domain.User, error)
	GetAllUsersMock     func() ([]domain.User, error)
}

func (u UserRepositoryMock) CreateAccount(user *domain.User) error {
	return u.CreateAccountMock(user)
}

func (u UserRepositoryMock) EmailExists(email string) (bool, error) {
	return u.EmailExistsMock(email)
}

func (u UserRepositoryMock) FindUserByEmail(email string) (*domain.User, error) {
	return u.FindUserByEmailMock(email)
}

func (u UserRepositoryMock) GetAllUsers() ([]domain.User, error) {
	return u.GetAllUsersMock()
}
