package repository

import (
	"first-task-alterra/internal/features/users"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string    `json:"name"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	BirthDate time.Time `json:"birth_date"`
	// Todos     []todos.Todo    `gorm:"foreignKey:Owner"`
}

func (u *User) toUserEntity() users.User {
	return users.User{
		ID:        u.ID,
		Name:      u.Name,
		Password:  u.Password,
		Email:     u.Email,
		Phone:     u.Phone,
		BirthDate: u.BirthDate,
	}
}

func toUserData(input users.User) User {
	return User{
		Name:      input.Name,
		Password:  input.Password,
		Email:     input.Email,
		Phone:     input.Phone,
		BirthDate: input.BirthDate,
	}
}
