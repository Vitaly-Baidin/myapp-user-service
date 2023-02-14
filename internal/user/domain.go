package user

import "github.com/google/uuid"

type User struct {
	Id       uuid.UUID
	Login    string
	Email    string
	Name     string
	Password string
}

func NewUser() *User {
	return &User{}
}
