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

func (us *UserRepository) Post(req *entities.User) (user *entities.User, err error) {
	user, err = us.orm.Post(req)
	if user.CreatedAt != req.CreatedAt {
		return user, exceptions.New(exceptions.AlreadyExists, "user " + user.UserId + " already exists")
	}
	return user, err
}

func (us *UserRepository) Patch(saldo float64, userId string) (user *entities.User, err error) {
	return nil, nil
}