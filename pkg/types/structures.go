package types

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type User struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	RoleId   string `json:"roleId"`
	Password string `json:"password"`
	MessId   string `json:"messId"`
	Email    string `json:"email"`
}

type BookRequest struct {
	ID          uint   `json:"bookID"`
	BookName    string `json:"bookname"`
	Author      string `json:"author"`
	Publication string `json:"publication,omitempty"`
}

func (book BookRequest) Validate() error {
	return validation.ValidateStruct(&book,
		validation.Field(&book.BookName,
			validation.Required.Error("Book name cannot be empty"),
			validation.Length(1, 50)),
		validation.Field(&book.Author,
			validation.Required.Error("Author name cannot be empty"),
			validation.Length(5, 50)),
	)
}

func (user User) ValidateUser() error {
	return validation.ValidateStruct(&user,
		validation.Field(&user.Name,
			validation.Required.Error("User name cannot be empty")),
		validation.Field(&user.RoleId,
			validation.Required.Error("Role Id cannot be empty")),
		validation.Field(&user.Password,
			validation.Required.Error("Password cannot be empty")),
		validation.Field(&user.MessId,
			validation.Required.Error("MessId cannot be empty")),
		validation.Field(&user.Email,
			validation.Required.Error("Email cannot be empty")),
	)
}
