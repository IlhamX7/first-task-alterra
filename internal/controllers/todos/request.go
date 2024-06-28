package todos

import (
	"first-task-alterra/internal/models"
)

type CreateRequest struct {
	Activity string `json:"activity"`
	UserId   int    `json:"user_id"`
}

type UpdateRequest struct {
	Activity string `json:"activity"`
	UserId   uint   `json:"user_id"`
	Mark     bool   `json:"mark"`
}

func ToModelTodosCreate(r CreateRequest) models.Todo {
	return models.Todo{
		Activity: r.Activity,
		Owner:    uint(r.UserId),
	}
}

func ToModelTodosUpdate(r UpdateRequest) models.Todo {
	return models.Todo{
		Activity: r.Activity,
		Owner:    r.UserId,
		Mark:     r.Mark,
	}
}
