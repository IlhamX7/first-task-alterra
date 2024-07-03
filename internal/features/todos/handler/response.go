package handler

import (
	"first-task-alterra/internal/features/todos"
)

type TodoResponse struct {
	ID       uint   `json:"id"`
	Activity string `json:"activity"`
	Mark     bool   `json:"mark"`
	Owner    uint   `json:"owner"`
}

func ToTodoReponse(input todos.Todo) TodoResponse {
	return TodoResponse{
		ID:       input.ID,
		Activity: input.Activity,
		Mark:     input.Mark,
		Owner:    input.Owner,
	}
}
