package user

import (
	"ms-user-api/exceptions"
	"ms-user-api/user/entities"
)

type UserRepository struct {
	orm entities.Storage
}

func NewUserRepository(orm entities.Storage) entities.Store {
	return &UserRepository{orm: orm}
}

func (us *UserRepository) List(userIds []string, limit int, page int) (user []entities.User, err error) {
	return us.orm.List(userIds, limit, page)
}

func (us *UserRepository) Post(req *entities.User) (user *entities.User, err error) {
	user, err = us.orm.Post(req)
	if user.CreatedAt != req.CreatedAt {
		return user, exceptions.New(exceptions.AlreadyExists, "user "+user.Id+" already exists")
	}
	return user, err
}

func (us *UserRepository) Patch(req *entities.User) (err error) {
	return us.orm.Patch(req)
}

func (us *UserRepository) Delete(req *entities.User) (err error) {
	return us.orm.Delete(req)
}
