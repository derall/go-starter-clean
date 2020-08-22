package entity

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" gorm:unique;type:varchar(30)`
	Email    string `json:"email" gorm:unique;type:varchar(30)`
	Password string `json:"password"`
}

type UserService interface {
	Store(*User) error
	FindAll() ([]*User, error)
}

type UserRepository interface {
	GetAll() ([]*User, error)
	Store(u *User) error
	Delete(id string) error
	GetByID(id string) (*User, error)
}
