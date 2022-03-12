package model

import (
	"main/internal/core/model/table"
	"main/internal/pkg/logger"
	"main/internal/pkg/objconv"
	"main/internal/pkg/pwd"
)

type Users []*User

type User struct {
	table.User
}

func NewUser(email, password, resource string) *User {
	u := User{
		table.User{
			Email:        email,
			Password:     []byte(password),
			AuthResource: resource,
		},
	}
	u.HashPassword()
	return &u
}

func (u *User) HashPassword() {
	hashed, err := pwd.Hash(u.Password)
	if err == nil {
		u.Password = hashed
	} else {
		logger.Logger.Errorf("error while hashing password: %v", err)
	}
}

func (u *User) GetID() uint {
	return u.ID
}

func (u *User) GetOwnerID() uint {
	return u.GetID()
}

func (u *User) ToMap() map[string]interface{} {
	return objconv.ToMap(u)
}

func (u *User) ToMapWithFields(fields []string) map[string]interface{} {
	return objconv.ToMapWithFields(u, fields)
}
