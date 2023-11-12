package user

import (
	"ms-user-api/exceptions"
	"ms-user-api/user/entities"
)

type UserRepository struct {
	orm entities.Storage
}

func NewUserRepository(orm entities.Storage) entities.Store {
	return &UserRepository{}
}

func (us *UserRepository) Get(userId string) (user *entities.User, err error) {
	user, err = us.orm.Get(userId)
	if err != nil {
		return nil, exceptions.New(exceptions.NotFound, "user with id " + userId + " not found")
	}
	return user, nil
}