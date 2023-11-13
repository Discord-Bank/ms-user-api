package router

import (
	"log"
	"ms-user-api/internal/http/health"
	userHandler "ms-user-api/internal/http/user"
	"ms-user-api/user/db"
	"os"

	"ms-user-api/user"

	"github.com/labstack/echo/v4"
)

func Handlers() (*echo.Echo){
	e := echo.New()

	orm, err := db.NewDatabase()
	if err != nil {
		log.Fatal("internal server error, fail to load database")
		os.Exit(1)
	}
	orm.AutoMigrateSetup()

	repo := user.NewUserRepository(orm)
	us := user.NewUserService(repo)
	uh := userHandler.NewUserHandler(us)

	h := health.NewHealthHandler()

	e.GET("/health", h.Health)

	v1 := e.Group("v1/users")
	v1.GET(":userId", uh.Get)
	v1.POST("", uh.Post)
	v1.PATCH(":userId", uh.Patch)
	v1.DELETE(":userId", uh.Delete)

	return e
}