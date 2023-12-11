package user

import (
	"errors"
	"ms-user-api/exceptions"
	"ms-user-api/user/entities"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	s entities.Service
}

func NewUserHandler(s entities.Service) *UserHandler {
	return &UserHandler{s: s}
}

func (uh *UserHandler) Post(c echo.Context) error {
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
					"user":    res,
				})
			}
			return uh.returnError(c, appErrors)
		}
	}

	return c.JSON(http.StatusCreated, res)
}

func (uh *UserHandler) List(c echo.Context) error {
	q := c.Request().URL.Query()
	ids := q["ids"]
	var page = 1
	pageParam, err := uh.getIntParam(c, "page")
	if err == nil {
		page = pageParam
	}
	var limit = 10
	limitParam, err := uh.getIntParam(c, "limit")
	if err == nil {
		limit = limitParam
	}

	response, err := uh.s.List(ids, limit, page)

	if err != nil {
		var appErrors *exceptions.Error
		if errors.As(err, &appErrors) {
			return uh.returnError(c, appErrors)
		} else {
			return uh.returnError(c, exceptions.New(exceptions.InternalServerError, "internal server error "+err.Error()))
		}
	}

	return c.JSON(http.StatusOK, response)
}

func (uh *UserHandler) Patch(c echo.Context) error {
	userId := c.Param("userId")
	var req *entities.UserPatchRequest
	if err := c.Bind(&req); err != nil {
		return uh.returnError(c, exceptions.New(exceptions.BadData, "unprocessable json"))
	}
	err := uh.s.Patch(req, userId)
	if err != nil {
		var appErrors *exceptions.Error
		if errors.As(err, &appErrors) {
			return uh.returnError(c, appErrors)
		}
	}

	return c.JSON(http.StatusNoContent, "")
}

func (uh *UserHandler) Delete(c echo.Context) error {
	userId := c.Param("userId")
	err := uh.s.Delete(userId)
	if err != nil {
		var appErrors *exceptions.Error
		if errors.As(err, &appErrors) {
			return uh.returnError(c, appErrors)
		}
	}

	return c.JSON(http.StatusNoContent, "")
}

func (h *UserHandler) printErrorMessage(c echo.Context, statusCode int, message string) error {
	return c.JSON(statusCode, map[string]string{
		"message": message,
	})
}

func (h *UserHandler) getIntParam(c echo.Context, paramName string) (int, error) {
	pageParam := c.QueryParam(paramName)
	if len(pageParam) > 0 {
		pageInt, err := strconv.Atoi(pageParam)
		if err == nil {
			return pageInt, nil
		}
	}
	return 0, errors.New("parameter not found")
}

func (h *UserHandler) returnError(c echo.Context, err *exceptions.Error) error {
	switch err.Code {
	case exceptions.NotFound:
		return h.printErrorMessage(c, http.StatusNotFound, err.Error())
	case exceptions.BadData:
		return h.printErrorMessage(c, http.StatusUnprocessableEntity, err.Error())
	case exceptions.BadRequest:
		return h.printErrorMessage(c, http.StatusBadRequest, err.Error())
	}

	return h.printErrorMessage(c, http.StatusInternalServerError, err.Error())
}
