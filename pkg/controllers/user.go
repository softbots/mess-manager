package controllers

import (
	"mess-manager/pkg/domain"
	"mess-manager/pkg/models"
	"mess-manager/pkg/types"
	"net/http"

	"github.com/labstack/echo/v4"
)

var UserService domain.IUserService

func SetUserService(userService domain.IUserService) {
	UserService = userService
}

func SignUp(e echo.Context) error {
	reqUser := &types.User{}

	if err := e.Bind(reqUser); err != nil {
		return e.JSON(http.StatusBadRequest, "Invalid user")
	}

	if err := reqUser.ValidateUser(); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	user := &models.User{
		Name:   reqUser.Name,
		RoleId: reqUser.RoleId,

		MessId:     reqUser.MessId,
		Email:      reqUser.Email,
		IsVerified: false,
	}

	if err := UserService.SignUp(user); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusCreated, "Book was created successfully")
}
