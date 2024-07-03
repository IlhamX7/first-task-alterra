package todos

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Todo struct {
	ID        uint
	Activity  string
	Mark      bool
	Owner     uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

// Error implements error.
func (t Todo) Error() string {
	panic("unimplemented")
}

type Handler interface {
	AddTodo() echo.HandlerFunc
	UpdateTodo() echo.HandlerFunc
	DeleteTodo() echo.HandlerFunc
	FindTodo() echo.HandlerFunc
}

type Services interface {
	AddTodo(newActivity string, userId uint) error
	UpdateTodo(id uint, newActivity string, newMark bool, owner uint) error
	DeleteTodo(id uint) (Todo, error)
	FindTodo(id uint) ([]Todo, error)
}

type Query interface {
	AddTodo(newTodo Todo) error
	UpdateTodo(updateTodo Todo) error
	DeleteTodo(id uint) (Todo, error)
	FindTodo(owner uint) ([]Todo, error)
	GetTodo(id uint) (Todo, error)
}
