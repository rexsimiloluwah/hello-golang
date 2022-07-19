package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/rexsimiloluwah/hello-golang/src/projects/golang-mongo-starter/internal/config"
	"github.com/rexsimiloluwah/hello-golang/src/projects/golang-mongo-starter/internal/service"
	usermock "github.com/rexsimiloluwah/hello-golang/src/projects/golang-mongo-starter/internal/tests/user"
	"github.com/rexsimiloluwah/hello-golang/src/projects/golang-mongo-starter/pkg/domain"
	"github.com/rexsimiloluwah/hello-golang/src/projects/golang-mongo-starter/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestRegisterAccount_UserExists(t *testing.T) {
	// test data
	testUser := `{"email":"rexsimiloluwa@gmail.com","password":"adetoyosi","username":"theblackdove"}`
	// echo setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(testUser))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// service setup
	cfg := &config.Settings{}
	mockUserDataStore := usermock.NewMockDataStore()
	userSvc := service.NewUserService(cfg, mockUserDataStore)

	mockApi := Api{
		userSvc: userSvc,
	}

	// act
	mockApi.Register(c)

	// assert
	response := models.Error{}
	_ = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, "USER_EXISTS", response.Name)
	assert.Equal(t, "Email already exists.", response.Message)
}

func TestRegisterAccount_UserExistsNo(t *testing.T) {
	// test data
	testUser := `{"email":"rexsimiloluwa2@gmail.com","password":"adetoyosi","username":"theblackdove"}`
	// echo setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(testUser))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// service setup
	cfg := &config.Settings{}
	mockUserDataStore := usermock.NewMockDataStore()
	userSvc := service.NewUserService(cfg, mockUserDataStore)

	mockApi := Api{
		userSvc: userSvc,
	}

	// act
	mockApi.Register(c)

	// assert
	assert.Equal(t, http.StatusCreated, rec.Code)
}

func TestGetAllUsers(t *testing.T) {
	// echo setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// service setup
	cfg := &config.Settings{}
	mockUserDataStore := usermock.NewMockDataStore()
	userSvc := service.NewUserService(cfg, mockUserDataStore)

	mockApi := Api{
		userSvc: userSvc,
	}

	// act
	mockApi.GetAllUsers(c)

	// assert
	response := []domain.User{}
	_ = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, 3, len(response))
	assert.Equal(t, "rexsimiloluwa@gmail.com", response[0].Email)
}

func TestLoginUser(t *testing.T) {
	// test data
	testUser := `{"email":"rexsimiloluwa2@gmail.com","password":"adetoyosi"}`
	// echo setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(testUser))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// service setup
	cfg := &config.Settings{
		JwtSecretKey: "secret",
		JwtExpiresIn: "24",
	}
	mockUserDataStore := usermock.NewMockDataStore()
	mockUserDataStore.CreateAccount(&domain.User{
		Email:    "rexsimiloluwa2@gmail.com",
		Password: "adetoyosi",
		Username: "theblackdove",
	})
	userSvc := service.NewUserService(cfg, mockUserDataStore)

	mockApi := Api{
		userSvc: userSvc,
	}

	// act
	mockApi.Login(c)

	// assert
	response := models.Response{}
	_ = json.Unmarshal(rec.Body.Bytes(), &response)
	fmt.Println(response)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, true, response.Status)
	assert.Equal(t, "Login successful", response.Message)
}
