package controllers

import (
	"mess-manager/pkg/domain"
	"mess-manager/pkg/models"
	"mess-manager/pkg/types"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

var BookService domain.IBookService

func SetBookService(bService domain.IBookService) {
	BookService = bService
}

func CreateBook(e echo.Context) error {
	reqBook := &types.BookRequest{}

	if err := e.Bind(reqBook); err != nil {
		return e.JSON(http.StatusBadRequest, "Invalid Data")
	}

	if err := reqBook.Validate(); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	book := &models.Book{
		BookName:    reqBook.BookName,
		Author:      reqBook.Author,
		Publication: reqBook.Publication,
	}

	if err := BookService.CreateBook(book); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusCreated, "Book was created successfully")
}

func GetBooks(e echo.Context) error {
	tempBookID := e.QueryParam("bookID")
	bookID, err := strconv.ParseInt(tempBookID, 0, 0)
	if err != nil && tempBookID != "" {
		return e.JSON(http.StatusBadRequest, "Enter a valid book ID")
	}

	book, err := BookService.GetBooks(uint(bookID))

	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	return e.JSON(http.StatusOK, book)
}
