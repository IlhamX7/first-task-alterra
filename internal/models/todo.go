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

func (tm *TodoModel) AddTodo(newActivity string, userId uint) (Todo, error) {
	var todo Todo
	todo.Activity = newActivity
	todo.Owner = userId
	err := tm.db.Create(&todo).Error
	if err != nil {
		return Todo{}, err
	}
	return todo, nil
}

func (tm *TodoModel) UpdateTodo(id uint, newActivity string, newMark bool, owner uint) (Todo, error) {
	var todo Todo
	err := tm.db.Where("ID = ?", id).First(&todo).Error
	if err != nil {
		return Todo{}, err
	}
	todo.Activity = newActivity
	todo.Mark = newMark
	todo.Owner = owner
	todo.UpdatedAt = time.Now()

	err = tm.db.Save(&todo).Error
	if err != nil {
		return Todo{}, err
	}

	return todo, nil
}

func (tm *TodoModel) DeleteTodo(id uint) (Todo, error) {
	var todo Todo
	err := tm.db.Where("ID = ?", id).First(&todo).Error
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
