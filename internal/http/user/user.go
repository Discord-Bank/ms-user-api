package user

import (
	"ms-user-api/user/entities"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler interface {
	Get(c echo.Context) (error)
}

type UserHandler struct {
	s entities.Service
}

func NewUserHandler(s entities.Service) Handler {
	return &UserHandler{s: s}
} 

func (uh *UserHandler) Get(c echo.Context) (error) {
	return c.JSON(http.StatusOK, map[string]string{"message":"first message"})
}