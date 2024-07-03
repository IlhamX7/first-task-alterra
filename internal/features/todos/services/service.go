package services

import (
	"first-task-alterra/internal/features/todos"
	"time"
)

type TodoServices struct {
	qry todos.Query
}

func NewTodoService(q todos.Query) todos.Services {
	return &TodoServices{
		qry: q,
	}
}

func (ts *TodoServices) AddTodo(newActivity string, userId uint) error {
	var todo todos.Todo
	todo.Activity = newActivity
	todo.Owner = userId
	todo.CreatedAt = time.Now()

	err := ts.qry.AddTodo(todo)
	if err != nil {
		return err
	}
	return nil
}

func (ts *TodoServices) UpdateTodo(id uint, newActivity string, newMark bool, owner uint) error {
	data, err := ts.qry.GetTodo(id)
	if err != nil {
		return err
	}
	data.Activity = newActivity
	data.Mark = newMark
	data.Owner = owner
	data.UpdatedAt = time.Now()

	err = ts.qry.UpdateTodo(data)
	if err != nil {
		return err
	}
	return nil
}
