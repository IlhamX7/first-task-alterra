package users

import (
	"first-task-alterra/internal/helper"
	"first-task-alterra/internal/models"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Name      string `json:"name"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	BirthDate string `json:"birth_date"` // example 2024-06-28
}

func ToModelUsers(r RegisterRequest) (models.User, error) {
	parsedTime, err := helper.StringToDate(r.BirthDate)
	if err != nil {
		return models.User{}, err
	}
	return models.User{
		Name:      r.Name,
		Password:  r.Password,
		Email:     r.Email,
		Phone:     r.Phone,
		BirthDate: parsedTime,
	}, nil
}
