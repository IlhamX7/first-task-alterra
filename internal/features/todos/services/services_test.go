package services_test

import (
	"first-task-alterra/internal/features/todos"
	mocksTodo "first-task-alterra/internal/features/todos/mocks"
	"first-task-alterra/internal/features/todos/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func FindTodo(t *testing.T) {
	qry := mocksTodo.NewQuery(t)
	srv := services.NewTodoService(qry)

	t.Run("Success Find", func(t *testing.T) {
		var num int = 1
		var uNum uint = uint(num)

		expectedTodos := []todos.Todo{
			{
				ID:       1,
				Activity: "Test Todo 1",
			},
			{
				ID:       2,
				Activity: "Test Todo 2",
			},
		}

		qry.On("FindTodo", uNum).Return(expectedTodos, nil)

		todos, err := srv.FindTodo(uNum)

		qry.AssertExpectations(t)

		assert.Nil(t, err)
		assert.Equal(t, expectedTodos, todos)
	})
}
