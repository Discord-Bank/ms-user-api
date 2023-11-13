package user

import (
	"ms-user-api/exceptions"
	"ms-user-api/user/entities"
	"time"
)

type UserService struct {
	repo entities.Store
}

func NewUserService(repo entities.Store) entities.Service {
	return &UserService{repo: repo}
}

func (us *UserService) Get(userId string) (*entities.User, error) {
	return us.repo.Get(userId)
}

func (us *UserService) Post(req *entities.UserRequest) (*entities.User, error) {
	res, err := us.toUser(req)
	if err != nil {
		return nil, err
	}

	return us.repo.Post(res)
}

func (us *UserService) Patch(req *entities.UserPatchRequest, userId string) (error) {
	if req.Saldo == nil {
		return exceptions.New(exceptions.BadRequest, "the field saldo is required")
	}
	return us.repo.Patch(&entities.User{UserId: userId, Saldo: *req.Saldo})
}

func (us *UserService) toUser(req *entities.UserRequest) (user *entities.User, err error) {
	if err = req.Validate(); err != nil {
		return user, exceptions.Wrap(exceptions.BadRequest, "invalid request body, error: ", err)
	}

	createdAt := time.Now()
	user = &entities.User{
		UserId: req.UserId,
		Saldo: 0,
		CreatedAt: &createdAt,
		IsActive: true,
	}

	return user, err
}