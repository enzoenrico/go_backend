package main

import (
	"fmt"

	// "github.com/enzoenrico/go_backend/app/database"

	"github.com/enzoenrico/go_backend/app/handlers"
	"github.com/labstack/echo/v4"
	// "github.com/labstack/echo/v4/middleware"
)

func main() {
	fmt.Print("> Starting the server")

	// check app/database/database.go for why the code is commented

	// db, err := database.GetDB("test_db")
	// if err != nil {
	// 	fmt.Println("> Error connecting to the database:%.*", err)
	// 	return
	// }
	// defer db.Close()

	e := echo.New()

	defer e.Close()

	// e.Use(middleware.CORS())
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set("Access-Control-Allow-Origin", "*")
			c.Response().Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			c.Response().Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			return next(c)
		}
	})

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if c.Request().Method == "OPTIONS" {
				return c.NoContent(204)
			}
			return next(c)
		}
	})

	// handler for the / route
	e.GET("/", func(c echo.Context) error {
		fmt.Printf("> Server starting on port 5000")
		return c.String(200, "ok")
	})

	// ========== USER ROUTES ==========
	// for the love of God please implement logs

	e.GET("/users", handlers.GetAllUsers)
	e.GET("/users/:id", handlers.GetUserByID)
	e.POST("/users", handlers.PostNewUser)

	// =========POST ROUTES =============
	e.GET("/all_posts", handlers.GetAllPosts)
	e.GET("/posts/:id", handlers.GetPostByID)
	e.POST("/posts", handlers.NewPost)

	//listen in port 5k and log it
	e.Logger.Fatal(e.Start(":5000"))
}
