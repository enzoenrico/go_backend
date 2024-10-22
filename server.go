package main

import (
	"fmt"

	"github.com/enzoenrico/go_backend/app/database"
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

    //TODO: implement a router   

    // ========== USER ROUTES ==========
	e.GET("/users", handlers.UserGetAllHandler)
	e.GET("/users/:id", handlers.UserGetHandler)
	e.POST("/users", handlers.UserPostHandler)



	//listen in port 5k and log it
	e.Logger.Fatal(e.Start(":5000"))
}
