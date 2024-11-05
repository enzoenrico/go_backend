package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/enzoenrico/go_backend/app"
	"github.com/enzoenrico/go_backend/app/handlers"
	"github.com/enzoenrico/go_backend/app/logger"
	"github.com/enzoenrico/go_backend/app/users"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"

	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

func main() {
	fmt.Println("> Starting the server")

	err := logger.InitLogger("./logs/logs.json")
	if err != nil {
		//funny how we log the logger not logging, lol
		fmt.Println("Erro ao inicializar o logger:", err)
		return
	}
	defer logger.Logger.Sync()

	logger.Logger.Info("Server is starting...")

	config, err := app.LoadConfig()
	if err != nil {
		logger.Logger.Error("Error loading config", zap.Error(err))
		return
	}

	// check app/database/database.go for why the code is commented
	// db, err := database.GetDB("test_db")
	// if err != nil {
	// 	fmt.Println("> Error connecting to the database:%.*", err)
	// 	return
	// }
	// defer db.Close()

	logger.Logger.Info("Config loaded successfully", zap.String("JWT Secret", config.JwtSecret))

	e := echo.New()
	defer e.Close()

	// CORS middleware
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set("Access-Control-Allow-Origin", "*")
			c.Response().Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			c.Response().Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			return next(c)
		}
	})

	// OPTIONS method handler
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if c.Request().Method == "OPTIONS" {
				return c.NoContent(204)
			}
			return next(c)
		}
	})

	// JWT-protected route

	userGroup := e.Group("/users", middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(config.JwtSecret),
	}))

	// Root route
	e.GET("/", func(c echo.Context) error {
		logger.Logger.Info("Root endpoint accessed")
		return c.File("./index.html")
	})

	// User routes
	userGroup.GET("", handlers.GetAllUsers)
	userGroup.GET("/:id", handlers.GetUserByID)
	userGroup.POST("", handlers.PostNewUser)

	// Post routes
	e.GET("/posts", handlers.GetAllPosts)
	e.GET("/posts/:id", handlers.GetPostByID)
	e.POST("/posts", handlers.NewPost)

	e.POST("/login", func(c echo.Context) error {
		var Iuser users.User
		c.Bind(Iuser)

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"name":      Iuser.Name,
			"timestamp": time.Now().Unix(),
		})
		tokenString, err := token.SignedString([]byte(config.JwtSecret))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{})
		}
		logger.Logger.Info("Accessed login route ", zap.String("token", tokenString))
		return c.JSON(http.StatusOK, map[string]string{"token": tokenString})
	})

	// Start server on port 5000
	logger.Logger.Info("Listening on port 5000")
	e.Logger.Fatal(e.Start(":5000"))
}
