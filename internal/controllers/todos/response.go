package todos

import (
	"first-task-alterra/internal/models"
)

type TodoResponse struct {
	ID       uint   `json:"id"`
	Activity string `json:"activity"`
	Mark     bool   `json:"mark"`
	Owner    uint   `json:"owner"`
}

func ToTodoReponse(input models.Todo) TodoResponse {
	return TodoResponse{
		ID:       input.ID,
		Activity: input.Activity,
		Mark:     input.Mark,
		Owner:    input.Owner,
	}
}
