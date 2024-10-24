package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/enzoenrico/go_backend/app/database"
	"github.com/enzoenrico/go_backend/app/users"
	"github.com/labstack/echo/v4"
)

func GetAllUsers(c echo.Context) error {
	jsonData, _ := json.MarshalIndent(database.UserDB, "", "    ")
	fmt.Printf("\r> Returning users: \n ")
	return c.String(http.StatusOK, string(jsonData))
}

func GetUserByID(c echo.Context) error {
	found_user, err := database.GetUser(c.Param("id"))
	if err != nil {
		// como implementamos um sistema robusto de logging?
		return c.JSON(http.StatusNotFound, "user not found")
	}
	return c.JSON(http.StatusOK, found_user)
}

func PostNewUser(c echo.Context) error {
	body_content := extract_user_from_request(c)
	new_user := users.User{
		Name:     body_content.Name,
		Email:    body_content.Email,
		Password: body_content.Password,
	}
	new_user_hash, err := new_user.GetHash()

	if err != nil {
		panic(err)
	}

	added, err := database.SetUser(new_user, new_user_hash)

	// fmt.Print(added)
	jsonData, _ := json.MarshalIndent(database.UserDB, "", "    ")
	fmt.Println("Updated user database: ")
	fmt.Println(string(jsonData))

	return c.JSON(http.StatusCreated, added)
}

func extract_user_from_request(c echo.Context) users.User {
	defer c.Request().Body.Close()

	b, _ := io.ReadAll(c.Request().Body)
	var user users.User

	if err := json.Unmarshal(b, &user); err != nil {
		panic(err)
	}
	return user
}
