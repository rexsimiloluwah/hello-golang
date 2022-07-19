package api

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rexsimiloluwah/hello-golang/src/projects/golang-mongo-starter/internal/config"
	"github.com/rexsimiloluwah/hello-golang/src/projects/golang-mongo-starter/internal/repository"
	"github.com/rexsimiloluwah/hello-golang/src/projects/golang-mongo-starter/internal/service"
	"github.com/rexsimiloluwah/hello-golang/src/projects/golang-mongo-starter/pkg/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type Api struct {
	server     *echo.Echo
	userSvc    service.IUserService
	cfg        *config.Settings
	middleware IMiddleware
}

func New(cfg *config.Settings, dbClient *mongo.Client) *Api {
	// New echo instance
	server := echo.New()

	// Add middlewares
	server.Use(middleware.Logger())
	server.Use(middleware.Recover())
	server.Use(middleware.RequestID())

	userRepository := repository.NewUserRepository(cfg, dbClient)
	userSvc := service.NewUserService(cfg, userRepository)

	return &Api{
		server:     server,
		userSvc:    userSvc,
		cfg:        cfg,
		middleware: NewMiddleware(cfg),
	}
}

func (a Api) ConfigureRoutes() {
	// Add routes
	a.server.GET("/api/v1/public/healthy", a.HealthCheck)
	a.server.POST("/api/v1/auth/register", a.Register)
	a.server.POST("/api/v1/auth/login", a.Login)
	a.server.GET("/api/v1/users", a.GetAllUsers)

	// protected routes
	protected := a.server.Group("api/v1")
	protected.Use(a.middleware.Auth)
	protected.GET("/secret", func(c echo.Context) error {
		userEmail := c.Get("user").(string)
		return c.String(http.StatusOK, userEmail)
	})
}

func (a Api) StartServer() {
	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8040"
	}

	a.ConfigureRoutes()
	// Start server
	a.server.Logger.Fatal(a.server.Start(fmt.Sprintf(":%s", PORT)))
}

// Register Handler
func (a Api) Register(c echo.Context) error {
	newUser, err := models.ValidateRegisterRequest(c)
	if err != nil {
		return c.JSON(err.Code, err)
	}

	err = a.userSvc.CreateAccount(newUser)
	if err != nil {
		return c.JSON(err.Code, err)
	}

	return c.String(http.StatusCreated, "")
}

// Login Handler
func (a Api) Login(c echo.Context) error {
	loginRequest, err := models.ValidateLoginRequest(c)
	if err != nil {
		return c.JSON(err.Code, err)
	}

	token, err := a.userSvc.Login(loginRequest)
	if err != nil {
		return c.JSON(err.Code, err)
	}
	response := models.Response{
		Status:  true,
		Message: "Login successful",
		Data: map[string]interface{}{
			"access_token": token,
		},
	}
	return c.JSON(http.StatusOK, response)
}

// Get all users Handler
func (a Api) GetAllUsers(c echo.Context) error {
	users, err := a.userSvc.FindAllUsers()
	if err != nil {
		fmt.Println(err)
		return c.JSON(err.Code, err)
	}
	return c.JSON(http.StatusOK, users)
}

// TODO: Health check Handler
func (a Api) HealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "Server is healthy!!")
}
