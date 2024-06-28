package users

import (
	"first-task-alterra/internal/models"
	"time"
)

type LoginResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	Phone     string    `json:"hp"`
	BirthDate time.Time `json:"birth_date"`
}

func ToLoginReponse(input models.User) LoginResponse {
	return LoginResponse{
		ID:        input.ID,
		Name:      input.Name,
		Email:     input.Email,
		Password:  input.Password,
		Phone:     input.Phone,
		BirthDate: input.BirthDate,
	}
}
