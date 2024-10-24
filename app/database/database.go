package database

// implementation not working for some reason
// that's what happens when u use your work's pc and don't have admin privileges
// lol

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/enzoenrico/go_backend/app/posts"
	"github.com/enzoenrico/go_backend/app/users"
	_ "github.com/mattn/go-sqlite3"
)

var UserDB = make(map[string]users.User)

var PostsDB = make(map[string]posts.Post)

func GetDB(dbName string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", strings.Join([]string{"./", dbName, ".db"}, ""))
	if err != nil {
		return nil, err
	}
	return db, nil
}

func CreateTable(db *sql.DB, tableName string, columns []string) (sql.Result, error) {
	query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s);", tableName, strings.Join(columns, ", "))
	res, err := db.Exec(query)
	return res, err
}
