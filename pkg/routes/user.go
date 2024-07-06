package routes

import (
	"mess-manager/pkg/controllers"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Echo) {
	user := e.Group("/users/")

	user.POST("register", controllers.SignUp)
}
