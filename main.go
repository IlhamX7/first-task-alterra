package main

import (
	"first-task-alterra/configs"
	"first-task-alterra/internal/controllers/todos"
	"first-task-alterra/internal/controllers/users"
	"first-task-alterra/internal/models"
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {
	cfg := configs.ImportSetting()
	db, err := configs.ConnectDB(cfg)
	if err != nil {
		fmt.Println("Stop program, masalah pada database", err.Error())
	}
	if err := db.AutoMigrate(&models.User{}, &models.Todo{}); err != nil {
		fmt.Println("Berhasil memasukan table user dan todo", err.Error())
	}

	um := models.NewUserModel(db)
	uc := users.NewUserController(um)

	tm := models.NewTodoModel(db)
	tc := todos.NewTodoController(tm)

	e := echo.New()
	e.POST("/register", uc.Register)
	e.POST("/login", uc.Login)

	e.POST("/todos", tc.AddTodo)
	e.PUT("/todos/:id", tc.UpdateTodo)
	e.GET("/todos/:id", tc.FindTodo)
	e.DELETE("/todos/:id", tc.DeleteTodo)

	e.Start(":8000")
}
