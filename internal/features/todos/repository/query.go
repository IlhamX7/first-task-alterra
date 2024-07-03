package repository

import (
	"errors"
	"first-task-alterra/internal/features/todos"
	"time"

	"gorm.io/gorm"
)

type TodoModel struct {
	db *gorm.DB
}

func NewTodoModel(connection *gorm.DB) todos.Query {
	return &TodoModel{
		db: connection,
	}
}

func (tm *TodoModel) AddTodo(todo todos.Todo) error {
	err := tm.db.Create(&todo).Error
	if err != nil {
		return err
	}
	return nil
}

func (tm *TodoModel) UpdateTodo(todo todos.Todo) error {
	err := tm.db.Save(&todo).Error
	if err != nil {
		return err
	}
	return nil
}

func (tm *TodoModel) GetTodo(id uint) (todos.Todo, error) {
	var todo todos.Todo
	err := tm.db.Where("ID = ?", id).First(&todo).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return todos.Todo{}, nil
		}
		return todos.Todo{}, err
	}
	return todo, nil
}

func (tm *TodoModel) FindTodo(owner uint) ([]todos.Todo, error) {
	var todo []todos.Todo
	err := tm.db.Where("owner = ?", owner).Find(&todo).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []todos.Todo{}, nil
		}
		return []todos.Todo{}, err
	}
	return todo, nil
}

func (tm *TodoModel) DeleteTodo(id uint) (todos.Todo, error) {
	var todo todos.Todo
	err := tm.db.Where("ID = ?", id).First(&todo).Error
	if err != nil {
		return todos.Todo{}, err
	}
	todo.DeletedAt = time.Now()
	err = tm.db.Save(&todo).Error
	if err != nil {
		return todos.Todo{}, err
	}

	return todo, nil
}
