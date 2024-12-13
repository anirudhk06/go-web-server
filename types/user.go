package types

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	CreateUser(User) error
	GetUserByID(ID int) (*User, error)
	FindUsers() ([]User, error)
}

type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"required,unique,not null"`
	Password string `json:"-" gorm:"required,not null"`
}

type RegisterPayload struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type LoginPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (p LoginPayload) Validate() map[string]string {
	errors := make(map[string]string)

	validate := validator.New()

	err := validate.Struct(p)

	if err != nil {
		for _, value := range err.(validator.ValidationErrors) {
			field := value.Field()
			tag := value.Tag()

			switch field {
			case "Email":
				if tag == "required" {
					errors["Email"] = "Please enter the email."
				} else if tag == "email" {
					errors["Email"] = "Please enter the valid email."
				}
			case "Password":
				if tag == "required" {
					errors["password"] = "Please enter the password."
				}
			}

		}
	}

	return errors
}
