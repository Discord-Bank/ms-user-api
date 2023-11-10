package user

import "ms-user-api/user/entities"

type UserRepository struct {}

func NewUserRepository() entities.Store {
	return &UserRepository{}
}