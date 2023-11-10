package user

import "ms-user-api/user/entities"

type UserService struct {
}

func NewUserService(repo entities.Store) entities.Service {
	return &UserService{}
}
