package user_test

import (
	"github.com/rexsimiloluwah/hello-golang/src/projects/golang-mongo-starter/internal/repository"
	"github.com/rexsimiloluwah/hello-golang/src/projects/golang-mongo-starter/pkg/domain"
)

type MockDataStore struct{}

var userMockData = []domain.User{
	{
		Email:    "rexsimiloluwa@gmail.com",
		Password: "adetoyosi",
		Username: "theblackdove",
	},
	{
		Email:    "adetoyosi2001@gmail.com",
		Password: "adetoyosi",
		Username: "acidic",
	},
}

func NewMockDataStore() repository.IUserRepository {
	return &MockDataStore{}
}

func (m *MockDataStore) CreateAccount(user *domain.User) error {
	userMockData = append(userMockData, *user)
	return nil
}

func (m *MockDataStore) EmailExists(email string) (bool, error) {
	for _, user := range userMockData {
		if user.Email == email {
			return true, nil
		}
	}
	return false, nil
}

func (m *MockDataStore) FindUserByEmail(email string) (*domain.User, error) {
	for _, user := range userMockData {
		if user.Email == email {
			return &user, nil
		}
	}
	return nil, nil
}

func (m *MockDataStore) GetAllUsers() ([]domain.User, error) {
	return userMockData, nil
}
