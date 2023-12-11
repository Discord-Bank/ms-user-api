package router

import (
	"ms-user-api/internal/http/health"
	userHandler "ms-user-api/internal/http/user"
	"ms-user-api/user/entities"

	"ms-user-api/user"

	"github.com/labstack/echo/v4"
)

func Handlers(orm entities.Storage) *echo.Echo {
	e := echo.New()

	repo := user.NewUserRepository(orm)
	us := user.NewUserService(repo)
	uh := userHandler.NewUserHandler(us)

	h := health.NewHealthHandler()

	e.GET("/health", h.Health)

	v1 := e.Group("v1/users")
	v1.GET("", uh.List)
	v1.POST("", uh.Post)
	v1.PATCH("/:userId", uh.Patch)
	v1.DELETE("/:userId", uh.Delete)

	return e
}
