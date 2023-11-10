package router

import (
	"ms-user-api/internal/http/health"
	userHandler "ms-user-api/internal/http/user"

	"ms-user-api/user"

	"github.com/labstack/echo/v4"
)

func Handlers() (*echo.Echo){
	e := echo.New()

	repo := user.NewUserRepository()
	us := user.NewUserService(repo)
	uh := userHandler.NewUserHandler(us)

	h := health.NewHealthHandler()

	e.GET("/health", h.Health)

	v1 := e.Group("v1/users")
	v1.GET(":userId", uh.Get)

	return e
}