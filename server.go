package main

import (
	"fmt"

	// "github.com/enzoenrico/go_backend/app/database"

	"github.com/enzoenrico/go_backend/app/handlers"
	"github.com/labstack/echo/v4"
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

	// handler for the / route
	e.GET("/", func(c echo.Context) error {
		// res, err := database.CreateTable(db, "teste", []string{"first", "second"})
		// if err != nil {
		// 	fmt.Println(err)
		// }
		// fmt.Println(res)
		return c.String(200, "ok")
	})

	// ========== USER ROUTES ==========
	//FIX: users returning base64
    // FIXED: actually retarded error, wasn't importing users into the db and the error didn't return a stck trace
    // for the love of God please implement logs
	e.GET("/all_users", handlers.GetAllUsers)
	e.GET("/users/:id", handlers.GetUserByID)
	e.POST("/users", handlers.PostNewUser)

	// =========POST ROUTES =============
	// e.GET("/posts/:id", handlers.GetPostByID)

	// database.PostsDB["first"] = posts.Post{
	// 	ID:        0,
	// 	Title:     "First Post",
	// 	Content:   "This is the first post.",
	// 	User:      users.User{},
	// 	Timestamp: 0,
	// }
	//listen in port 5k and log it
	e.Logger.Fatal(e.Start(":5000"))
}
