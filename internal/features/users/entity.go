package users

import (
	"time"

	"github.com/labstack/echo/v4"
)

type User struct {
	ID        uint
	Name      string
	Password  string
	Email     string
	Phone     string
	BirthDate time.Time
	CreatedAt time.Time
}

type Handler interface {
	Register() echo.HandlerFunc
	Login() echo.HandlerFunc
}

type Services interface {
	Register(newUser User) error
	Login(email string, password string) (User, string, error)
}

type Query interface {
	Register(newUser User) error
	Login(email string) (User, error)
}
