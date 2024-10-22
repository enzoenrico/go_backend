package users

import (
	"golang.org/x/crypto/bcrypt"
	"strings"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func New(name, email, password string) User {
	u := User{}
	u.Name = name
	u.Email = email
	u.Password = password
	return u
}

func (u *User) gen_sum_str() string {
	hashable := strings.Join([]string{u.Name, u.Email, u.Password}, "-")
	return hashable
}

func (u *User) GetHash() (string, error) {

	bytes, err := bcrypt.GenerateFromPassword([]byte(u.gen_sum_str()), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
