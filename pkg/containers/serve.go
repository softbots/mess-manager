package containers

import (
	"fmt"
	"log"
	"mess-manager/pkg/config"
	"mess-manager/pkg/connection"
	"mess-manager/pkg/controllers"
	"mess-manager/pkg/repositories"
	"mess-manager/pkg/routes"
	"mess-manager/pkg/services"

	"github.com/labstack/echo/v4"
)

func Serve(e *echo.Echo) {

	config.SetConfig()
	// config initialization

	db := connection.GetDB()

	userRepo := repositories.UserDbInstance(db)

	bookRepo := repositories.BookDBInstance(db)

	bookService := services.BookServiceInstance(bookRepo)
	userService := services.UserServiceInstance(userRepo)

	controllers.SetBookService(bookService)
	controllers.SetUserService(userService)

	routes.BookRoutes(e)
	routes.UserRoutes(e)

	log.Fatal(e.Start(fmt.Sprintf(":%s", config.LocalConfig.PORT)))
}
