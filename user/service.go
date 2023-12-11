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

func (us *UserService) List(userIds []string, limit int, page int) ([]*entities.UserResponse, error) {
	usr, err := us.repo.List(userIds, limit, page)
	user := []*entities.UserResponse{}
	for _, u := range usr {
		user = append(user, us.toUserResponse(&u))
	}
	return user, err
}

func (us *UserService) Post(req *entities.UserRequest) (*entities.UserResponse, error) {
	res, err := us.toUser(req)
	if err != nil {
		return nil, err
	}
	res, err = us.repo.Post(res)
	resp := us.toUserResponse(res)
	return resp, err
}

func (us *UserService) Patch(req *entities.UserPatchRequest, userId string) error {
	if req.Saldo == nil {
		return exceptions.New(exceptions.BadRequest, "the field saldo is required")
	}
	return us.repo.Patch(&entities.User{Id: userId, Saldo: *req.Saldo})
}

func (us *UserService) Delete(userId string) error {
	return us.repo.Delete(&entities.User{Id: userId})
}

func (us *UserService) toUserResponse(user *entities.User) *entities.UserResponse {
	return &entities.UserResponse{
		Id:        user.Id,
		CreatedAt: user.CreatedAt,
		Saldo:     user.Saldo,
	}
}

func (us *UserService) toUser(req *entities.UserRequest) (user *entities.User, err error) {
	if err = req.Validate(); err != nil {
		return user, exceptions.Wrap(exceptions.BadRequest, "invalid request body, error: ", err)
	}

	createdAt := time.Now()
	user = &entities.User{
		Id:        req.Id,
		Saldo:     0,
		CreatedAt: createdAt,
	}

	return user, err
}
