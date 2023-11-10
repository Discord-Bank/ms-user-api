package entities

import "time"

type Store interface {}

type Service interface {}

type User struct {
	UserId string `json:"userId"`
	Saldo float64 `json:"saldo"`
	CreatedAt *time.Time `json:"createdAt"`
	IsActive bool `json:"isActive"`
}

type UserPatchRequest struct {
	IsActive *bool `json:"isActive"`
}