package entities

import (
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Storage interface {
	List(userIds []string, limit int, page int) (user []User, err error)
	Post(req *User) (*User, error)
	Patch(req *User) (err error)
	Delete(req *User) (err error)
	AutoMigrateSetup()
}

type Store interface {
	List(userIds []string, limit int, page int) (user []User, err error)
	Post(req *User) (user *User, err error)
	Delete(req *User) (err error)
	Patch(req *User) (err error)
}

type Service interface {
	List(userIds []string, limit int, page int) (user []*UserResponse, err error)
	Post(req *UserRequest) (user *UserResponse, err error)
	Patch(req *UserPatchRequest, userId string) (err error)
	Delete(userId string) error
}

type User struct {
	Id        string `gorm:"primaryKey"`
	Saldo     float64
	CreatedAt time.Time
	gorm.DeletedAt
}

type UserRequest struct {
	Id string `json:"id" validate:"required,min=3"`
}

type UserResponse struct {
	Id        string    `json:"id"`
	Saldo     float64   `json:"saldo"`
	CreatedAt time.Time `json:"createdAt"`
}

func (u *UserRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}

type UserPatchRequest struct {
	Saldo *float64 `json:"saldo"`
}
