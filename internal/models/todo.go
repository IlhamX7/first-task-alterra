package models

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Activity string
	Mark     bool
	Owner    uint
}

type TodoModel struct {
	db *gorm.DB
}

func NewTodoModel(connection *gorm.DB) *TodoModel {
	return &TodoModel{
		db: connection,
	}
}

func (tm *TodoModel) AddTodo(newData Todo) (Todo, error) {
	newData.Mark = false
	err := tm.db.Create(&newData).Error
	if err != nil {
		return Todo{}, err
	}
	return newData, nil
}

func (tm *TodoModel) UpdateTodo(owner uint, newActivity string) (Todo, error) {
	var todo Todo
	err := tm.db.Where("owner = ?", owner).First(&todo).Error
	if err != nil {
		return Todo{}, err
	}
	todo.Activity = newActivity
	todo.UpdatedAt = time.Now()

	err = tm.db.Save(&todo).Error
	if err != nil {
		return Todo{}, err
	}

	return todo, nil
}

func (tm *TodoModel) DeleteTodo(owner uint) (Todo, error) {
	var todo Todo
	err := tm.db.Where("owner = ?", owner).First(&todo).Error
	if err != nil {
		return Todo{}, err
	}
	err = tm.db.Delete(&todo).Error
	if err != nil {
		return Todo{}, err
	}

	return todo, nil
}

func (tm *TodoModel) FindTodo(owner uint) ([]Todo, error) {
	var todo []Todo
	err := tm.db.Where("owner = ?", owner).Find(&todo).Error
	if err != nil {
		return nil, err
	}

	return todo, nil
}
