package models

import (
	"fmt"
	"strings"

	"github.com/anirudhk06/go-web-server/configs"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Email    string `gorm:"unique"`
	Password string
}

type RegisterPayload struct {
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"min=8"`
	ConfirmPassword string `json:"confirm_password" validate:"eqfield=Password"`
}

func (p RegisterPayload) Validate() map[string]string {
	p.Email = strings.ToLower(strings.TrimSpace(p.Email))
	p.Password = strings.ReplaceAll(strings.TrimSpace(p.Password), " ", "")
	p.ConfirmPassword = strings.TrimSpace(p.ConfirmPassword)
	validate := validator.New()
	err := validate.Struct(p)

	result := map[string]string{}

	isExists := configs.DB.Where("email = ?", p.Email).First(&User{}).Error

	if isExists == nil {
		result["email"] = fmt.Sprintf("User with email '%s' already exists", p.Email)
		return result
	}

	if err != nil {
		errs := err.(validator.ValidationErrors)

		for _, e := range errs {
			field := e.Field()
			tag := e.Tag()

			switch field {
			case "Email":
				if tag == "email" {
					result["email"] = "Please enter valid email"
				} else if tag == "required" {
					result["email"] = "Please enter the email"
				}
			case "Password":
				if tag == "required" {
					result["password"] = "Please enter the password"
				} else if tag == "min" {
					result["password"] = "Please must contains 8 charactors"
				}
			case "ConfirmPassword":
				if tag == "required" {
					result["confirm_password"] = "Please enter confirm the password"
				} else if tag == "eqfield" {
					result["confirm_password"] = "Password not matching"
				}
			}
		}
	}

	return result
}
