package app

import (
	"github.com/banggibima/go-echo-restful-api/internal/transport/rest"
	"github.com/labstack/echo/v4"
)

type App struct {
	UserHandler *rest.UserHandler
}

func NewApp(userHandler *rest.UserHandler) *App {
	return &App{UserHandler: userHandler}
}

func (a *App) SetupRoutes(e *echo.Echo) {
	v1 := e.Group("/v1")
	{
		users := v1.Group("/users")
		{
			users.GET("/", a.UserHandler.GetUsersHandler)
			users.GET("/:id", a.UserHandler.GetUserByIDHandler)
			users.POST("/", a.UserHandler.CreateUserHandler)
			users.PUT("/:id", a.UserHandler.UpdateUserHandler)
			users.DELETE("/:id", a.UserHandler.DeleteUserHandler)
		}
	}
}
