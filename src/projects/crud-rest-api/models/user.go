package models

import (
	"strings"
	"time"

	u "../utils"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID          uint       `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Name        string     `gorm:"type:varchar(200);not null" json:"name"`
	Email       string     `gorm:"type:varchar(300);not null;unique_index" json:"email"`
	Password    string     `gorm:"type:varchar(100);not null" json:"password"`
	Bio         string     `gorm:"type:varchar(500)" json:"bio,omitempty"`
	Phone       string     `gorm:"type:varchar(100)" json:"phone"`
	Picture     string     `gorm:"type:varchar(300)" json:"picture,omitempty"`
	Status      string     `gorm:"type:varchar(500)" json:"status"`
	Website     string     `gorm:"type:varchar(300)" json:"website"`
	AccessToken string     `json:"access_token,omitempty";sql:"-"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

// Validation
// Returns message, ok status
func (user *User) Validate() (map[string]interface{}, bool) {
	if !strings.Contains(user.Email, "@") {
		return u.Message(false, "Email is required."), false
	}

	if len(user.Password) < 8 {
		return u.Message(false, "Password must be more than 8 characters."), false
	}

	// To check for duplicate emails
	existingUser := &User{}
	err := Db().Table("users").Where("email = ?", user.Email).First(existingUser)
	if err != nil {
		if err.Error != gorm.ErrRecordNotFound {
			return u.Message(false, "A user with this mail already exists."), false
		}
	}

	return u.Message(false, "Validation successful."), true
}

// Create a new user
func (user *User) CreateUser() map[string]interface{} {
	resp, ok := user.Validate()
	if !ok {
		return resp
	}

	// Generate hashed password using bcrypt
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)
	// Add the new user to database easily with Gorm DB
	Db().Create(user)

	if user.ID <= 0 {
		return u.Message(false, "Failed to create user, connection error.")
	}

	response := u.Message(true, "User successfully created.")
	response["user"] = user
	return response
}

// Login an existing user
// The Login method returns a signed JSON web token for the user
func Login(email string, password string) map[string]interface{} {
	user := &User{}
	err := Db().Table("users").Where("email = ?", email).First(user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(false, "User does not exist.")
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		// When the two passwords do not match
		return u.Message(false, "Incorrect password.")
	}

	user.Password = ""
	// Create JWT token
	token, err := CreateToken(user)
	user.AccessToken = token
	response := u.Message(true, "User logged in successfully.")
	response["user"] = user
	return response
}

// To get user from database by ID
func GetUser(userId uint) map[string]interface{} {
	user := &User{}
	var response map[string]interface{}
	err := Db().Table("users").Where("id = ?", userId).First(user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			response = u.Message(false, "User not found.")
			response["user"] = nil
			return response
		}
	}

	user.Password = ""
	response = u.Message(true, "Successfully fetched user.")
	response["user"] = user
	return response
}
