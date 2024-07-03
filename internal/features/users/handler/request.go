package handler

import (
	"first-task-alterra/internal/features/users"
	"first-task-alterra/internal/utils"
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

func ToModelUsers(r RegisterRequest) users.User {
	parsedTime, err := utils.StringToDate(r.BirthDate)
	if err != nil {
		return users.User{}
	}
	return users.User{
		Name:      r.Name,
		Password:  r.Password,
		Email:     r.Email,
		Phone:     r.Phone,
		BirthDate: parsedTime,
	}
}
