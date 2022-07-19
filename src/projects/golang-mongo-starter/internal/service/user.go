package service

import (
	"fmt"
	"log"

	"github.com/rexsimiloluwah/hello-golang/src/projects/golang-mongo-starter/internal/config"
	"github.com/rexsimiloluwah/hello-golang/src/projects/golang-mongo-starter/internal/repository"
	"github.com/rexsimiloluwah/hello-golang/src/projects/golang-mongo-starter/internal/utils"
	"github.com/rexsimiloluwah/hello-golang/src/projects/golang-mongo-starter/pkg/domain"
	"github.com/rexsimiloluwah/hello-golang/src/projects/golang-mongo-starter/pkg/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IUserService interface {
	CreateAccount(user *domain.User) *models.Error
	Login(user *domain.User) (string, *models.Error)
	FindAllUsers() ([]domain.User, *models.Error)
}

type UserService struct {
	userRepository repository.IUserRepository
	cfg            *config.Settings
	jwt            JwtServiceInterface
}

func NewUserService(cfg *config.Settings, userRepository repository.IUserRepository) IUserService {
	return &UserService{
		userRepository: userRepository,
		cfg:            cfg,
		jwt:            NewJwtService(cfg),
	}
}

func (u UserService) CreateAccount(user *domain.User) *models.Error {
	userExists, err := u.userRepository.EmailExists(user.Email)
	if err != nil {
		return &models.Error{
			Code:    500,
			Name:    "SERVER_ERROR",
			Message: "Something went wrong",
			Error:   err.Error(),
		}
	}

	if userExists {
		return &models.Error{
			Code:    400,
			Name:    "USER_EXISTS",
			Message: "Email already exists.",
		}
	}

	user.ID = primitive.NewObjectID()
	hash, err := utils.HashPassword(user.Password)

	if err != nil {
		return &models.Error{
			Code:    500,
			Name:    "SERVER_ERROR",
			Message: "Error hashing password.",
			Error:   err.Error(),
		}
	}

	user.Password = hash
	err = u.userRepository.CreateAccount(user)

	if err != nil {
		return &models.Error{
			Code:    500,
			Name:    "SERVER_ERROR",
			Message: "Error creating user.",
			Error:   err.Error(),
		}
	}

	return nil
}

func (u UserService) Login(user *domain.User) (string, *models.Error) {
	// check if the user exists
	existingUser, err := u.userRepository.FindUserByEmail(user.Email)
	log.Println(existingUser)
	if err != nil {
		return "", &models.Error{
			Code:    500,
			Name:    "SERVER_ERROR",
			Message: "An error occurred.",
			Error:   err.Error(),
		}
	}

	if existingUser == nil {
		return "", &models.Error{
			Code:    400,
			Name:    "INVALID_LOGIN",
			Message: "User email does not exist.",
		}
	}

	// compare hashed passwords
	err = utils.ComparePasswordHash(user.Password, existingUser.Password)

	if err != nil {
		return "", &models.Error{
			Code:    400,
			Name:    "INVALID_LOGIN",
			Message: "Password is incorrect.",
			Error:   err.Error(),
		}
	}

	// generate the JWT access token for the user
	token, err := u.jwt.GenerateToken(user.Email, user.ID.String())
	if err != nil {
		return "", &models.Error{
			Code:    500,
			Name:    "SERVER_ERROR",
			Message: "Error generating JWT token.",
			Error:   err.Error(),
		}
	}

	return token, nil
}

func (u UserService) FindAllUsers() ([]domain.User, *models.Error) {
	users, err := u.userRepository.GetAllUsers()
	fmt.Println(err)
	if err != nil {
		return []domain.User{}, &models.Error{
			Code:    500,
			Name:    "SERVER_ERROR",
			Message: "Error finding all users.",
			Error:   err.Error(),
		}
	}

	return users, nil
}
