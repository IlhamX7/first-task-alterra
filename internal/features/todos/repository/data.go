package repository

import (
	"first-task-alterra/internal/features/todos"

	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Activity string `json:"activity"`
	Mark     bool   `json:"mark"`
	Owner    uint   `json:"owner"`
}

func (u *Todo) toTodoEntity() todos.Todo {
	return todos.Todo{
		ID:       u.ID,
		Activity: u.Activity,
		Mark:     u.Mark,
		Owner:    u.Owner,
	}
}

func toTodoData(input todos.Todo) Todo {
	return Todo{
		Activity: input.Activity,
		Mark:     input.Mark,
		Owner:    input.Owner,
	}
}
