package entities

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Storage interface {
	Get(userId string) (user *User, err error)
	Post(req *User) (*User, error)
	AutoMigrateSetup()
}

type Store interface {
	Get(userId string) (user *User, err error)
	Post(req *User) (user *User, err error)
	Patch(attr map[string]interface{}, userId string) (user *User, err error)
}

type Service interface {
	Get(userId string) (user *User, err error)
	Post(req *UserRequest) (user *User, err error)
	Patch(req *UserPatchRequest, userId string) (user *User, err error)
}

type User struct {
	UserId string `json:"userId"`
	Saldo float64 `json:"saldo"`
	CreatedAt *time.Time `json:"createdAt"`
	IsActive bool `json:"isActive"`
}

type UserRequest struct {
	UserId string `json:"userId" validate:"required,min=3"`
}


func (u *UserRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}

type UserPatchRequest struct {
	Saldo *float64 `json:"saldo"`
}