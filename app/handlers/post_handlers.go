package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/enzoenrico/go_backend/app/database"
	"github.com/enzoenrico/go_backend/app/posts"
	"github.com/labstack/echo/v4"
)

func GetPostByID(c echo.Context) error {
	id := c.Param("id")
	fmt.Println("ID: ", id)
	return c.JSON(http.StatusFound, database.PostsDB[id])
}

func NewPost(c echo.Context) error {
	body_content := extract_post_from_request(c)
	new_post := posts.Post{
		ID:        body_content.ID,
		Title:     body_content.Title,
		Content:   body_content.Content,
		User:      body_content.User,
		Timestamp: body_content.Timestamp,
	}
	database.PostsDB[strconv.Itoa(new_post.ID)] = new_post
	return c.JSON(http.StatusCreated, body_content)
}

func extract_post_from_request(c echo.Context) posts.Post {
	defer c.Request().Body.Close()

	b, _ := io.ReadAll(c.Request().Body)
	var post posts.Post
	if err := json.Unmarshal(b, &post); err != nil {
		panic(err)
	}
	return post
}
