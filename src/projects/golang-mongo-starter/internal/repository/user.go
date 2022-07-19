package repository

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/rexsimiloluwah/hello-golang/src/projects/golang-mongo-starter/internal/config"
	"github.com/rexsimiloluwah/hello-golang/src/projects/golang-mongo-starter/pkg/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IUserRepository interface {
	CreateAccount(user *domain.User) error
	EmailExists(email string) (bool, error)
	FindUserByEmail(email string) (*domain.User, error)
	GetAllUsers() ([]domain.User, error)
}

type UserRepository struct {
	userCollection *mongo.Collection
	ctx            context.Context
}

func NewUserRepository(cfg *config.Settings, dbClient *mongo.Client) IUserRepository {
	userCollection := dbClient.Database(cfg.DbName).Collection("users")
	return &UserRepository{
		userCollection: userCollection,
		ctx:            context.TODO(),
	}
}

func (u UserRepository) CreateAccount(user *domain.User) error {
	_, err := u.userCollection.InsertOne(u.ctx, user)
	if err != nil {
		return errors.Wrap(err, "Error inserting new user.")
	}
	return nil
}

func (u UserRepository) EmailExists(email string) (bool, error) {
	var existingUser *domain.User
	filter := bson.D{primitive.E{Key: "email", Value: email}}

	if err := u.userCollection.FindOne(u.ctx, filter).Decode(&existingUser); err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, errors.Wrap(err, "Error finding user by email")
	}
	fmt.Println("existing user", existingUser)
	return true, nil
}

func (u UserRepository) FindUserByEmail(email string) (*domain.User, error) {
	var existingUser domain.User
	filter := bson.D{primitive.E{Key: "email", Value: email}}

	if err := u.userCollection.FindOne(u.ctx, filter).Decode(&existingUser); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, errors.Wrap(err, "Error finding user by email")
	}
	fmt.Println("existing user", existingUser)
	return &existingUser, nil
}

func (u UserRepository) GetAllUsers() ([]domain.User, error) {
	var users []domain.User
	filter := bson.D{}
	projection := bson.D{{"username", 1}, {"email", 1}, {"_id", 1}}
	opts := options.Find().SetProjection(projection)

	cursor, err := u.userCollection.Find(u.ctx, filter, opts)
	if err != nil {
		return nil, err
	}

	if err = cursor.All(u.ctx, &users); err != nil {
		return nil, err
	}
	return users, nil
}
