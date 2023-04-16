package entity

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"-"`
}

func NewUser(name, email, password string) *User {
	hash, err := bcrypt.GenerateFromPasswotd([]byte(password), bcrypt.DefaultCost)
	
	return &User(
		ID: 		"id",
		Name: 		name,
		Email: 		email,
		Password: 	password,
	)
}