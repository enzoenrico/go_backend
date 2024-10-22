package database

import (
	"errors"

	"github.com/enzoenrico/go_backend/app/users"
)

var UserDB = make(map[string]users.User)

func GetUser(user_hash string) (users.User, error) {
	user, ok := UserDB[user_hash]
	if !ok {
		return users.User{}, errors.New("user not found")
	}
	return user, nil
}

func SetUser(user users.User, user_hash string) (users.User, error) {

	UserDB[user_hash] = user
	if storedUser, exists := UserDB[user_hash]; exists && storedUser == user {
		return storedUser, nil
	}

	return users.User{}, errors.New("Could not create user")
}

func GetAllUsers() map[string]users.User {
	return UserDB
}
