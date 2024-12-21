package types

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	CreateUser(User) error
	GetUserByID(ID int) (*User, error)
	FindUsers(page, limit int) ([]User, int64, error)
}

type User struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	Email     string    `json:"email" gorm:"required,unique,not null"`
	Password  string    `json:"-" gorm:"required,not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
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
