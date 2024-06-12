package controllers

import (
	"first-task-alterra/internal/models"
	"fmt"
)

type TodoController struct {
	model *models.TodoModel
}

func NewTodoController(m *models.TodoModel) *TodoController {
	return &TodoController{
		model: m,
	}
}

func (tc *TodoController) AddTodo(id uint) (bool, error) {
	var newData models.Todo
	fmt.Print("Masukkan Aktivitas ")
	fmt.Scanln(&newData.Activity)
	newData.Owner = id
	_, err := tc.model.AddTodo(newData)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (tc *TodoController) UpdateTodo(id uint) (bool, error) {
	var updateData models.Todo
	fmt.Println("Update Aktivitas ")
	fmt.Scanln(&updateData.Activity)
	_, err := tc.model.UpdateTodo(id, updateData.Activity)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (tc *TodoController) DeleteTodo(id uint) (bool, error) {
	_, err := tc.model.DeleteTodo(id)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (tc *TodoController) FindTodo(id uint) ([]map[string]any, error) {
	data, err := tc.model.FindTodo(id)
	if err != nil {
		return nil, err
	}

	var result []map[string]any
	for _, todo := range data {
		todoMap := map[string]any{
			"id":       todo.ID,
			"activity": todo.Activity,
			"owner":    todo.Owner,
		}
		result = append(result, todoMap)
	}

	return result, nil
}
