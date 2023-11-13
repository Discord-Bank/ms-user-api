package user

import (
	"errors"
	"ms-user-api/exceptions"
	"ms-user-api/user/entities"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler interface {
	Get(c echo.Context) (error)
	Post(c echo.Context) (error)
	Patch(c echo.Context) (error)
	Delete(c echo.Context) (error)
}

type UserHandler struct {
	s entities.Service
}

func NewUserHandler(s entities.Service) Handler {
	return &UserHandler{s: s}
} 

func (uh *UserHandler) Post(c echo.Context) (error) {
	var req *entities.UserRequest
	if err := c.Bind(&req); err != nil {
		return uh.returnError(c, exceptions.New(exceptions.BadData, "unprocessable json"))
	}

	res, err := uh.s.Post(req)
	if err != nil {
		var appErrors *exceptions.Error
		if errors.As(err, &appErrors) {
			if appErrors.Code == exceptions.AlreadyExists {
				return c.JSON(http.StatusNotImplemented, map[string]interface{}{
					"message": err.Error(),
					"user": res,
				})
			}
			return uh.returnError(c, appErrors)
		}
	}

	return c.JSON(http.StatusCreated, res)
}

func (uh *UserHandler) Get(c echo.Context) (error) {
	userId := c.Param("userId")
	
	user, err := uh.s.Get(userId)
	if err != nil {
		var appErrors *exceptions.Error
		if errors.As(err, &appErrors){
			return uh.returnError(c, appErrors)
		}
	}

	return c.JSON(http.StatusOK, user)
}

func (uh *UserHandler) Patch(c echo.Context) (error) {
	return c.JSON(http.StatusNoContent, "")
}

func (uh *UserHandler) Delete(c echo.Context) (error) {
	return c.JSON(http.StatusOK, map[string]string{"message": "deleted"})
}

func (h *UserHandler) printErrorMessage(c echo.Context, statusCode int, message string) error {
	return c.JSON(statusCode, map[string]string{
		"message": message,
	})
}

func (h *UserHandler) returnError(c echo.Context, err *exceptions.Error) error {
	switch err.Code {
	case exceptions.NotFound:
		return h.printErrorMessage(c, http.StatusNotFound, err.Message)
	case exceptions.BadData:
		return h.printErrorMessage(c, http.StatusUnprocessableEntity, err.Message)
	case exceptions.BadRequest:
		return h.printErrorMessage(c, http.StatusBadRequest, err.Message)
	}

	return h.printErrorMessage(c, http.StatusInternalServerError, err.Message)
}