package entities

import "time"

type Storage interface {
	Get(userId string) (user *User, err error)
	AutoMigrateSetup()
}

type Store interface {
	Get(userId string) (user *User, err error)
}

type Service interface {
	Get(userId string) (user *User, err error)
	Post(req *UserRequest) (user *User, err error)
}

type User struct {
	UserId string `json:"userId"`
	Saldo float64 `json:"saldo"`
	CreatedAt *time.Time `json:"createdAt"`
	IsActive bool `json:"isActive"`
}

type UserRequest struct {
	UserId string `json:"userId"`
}

type UserPatchRequest struct {
	IsActive *bool `json:"isActive"`
}