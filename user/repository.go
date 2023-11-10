package user

import "ms-user-api/user/entities"

type UserRepository struct {}

func NewUserRepository() entities.Store {
	return &UserRepository{}
}

func (us *UserRepository) Get(userId string) (user *entities.User, err error) {
	return nil, nil
}