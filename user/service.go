package user

import "ms-user-api/user/entities"

type UserService struct {
	repo entities.Store
}

func NewUserService(repo entities.Store) entities.Service {
	return &UserService{repo: repo}
}

func (us *UserService) Get(userId string) (user *entities.User, err error) {
	return us.repo.Get(userId)
}