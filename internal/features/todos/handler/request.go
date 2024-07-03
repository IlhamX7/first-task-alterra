package handler

import (
	"first-task-alterra/internal/features/todos"
)

type CreateRequest struct {
	Activity string `json:"activity"`
}

type UpdateRequest struct {
	Activity string `json:"activity"`
	UserId   uint   `json:"user_id"`
	Mark     bool   `json:"mark"`
}

func ToModelTodosCreate(r CreateRequest) todos.Todo {
	return todos.Todo{
		Activity: r.Activity,
	}
}

func ToModelTodosUpdate(r UpdateRequest) todos.Todo {
	return todos.Todo{
		Activity: r.Activity,
		Owner:    r.UserId,
		Mark:     r.Mark,
	}
}
