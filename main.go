package main

import (
	"first-task-alterra/configs"
	"first-task-alterra/internal/controllers/todos"
	"first-task-alterra/internal/controllers/users"
	"first-task-alterra/internal/models"
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
	e.GET("/hello", func(c echo.Context) error {
		return c.JSON(200, "hello world")
	})

	e.POST("/register", uc.Register)
	e.POST("/login", uc.Login)

	jwtKey := os.Getenv("JWT_SECRET")
	if jwtKey == "" {
		fmt.Println("JWT secret key not found in environment variables")
	}

	t := e.Group("/todos")
	t.Use(echojwt.WithConfig(
		echojwt.Config{
			SigningKey:    []byte(jwtKey),
			SigningMethod: jwt.SigningMethodHS256.Name,
		},
	))
	t.POST("", tc.AddTodo)
	t.PUT("/:id", tc.UpdateTodo)
	t.GET("", tc.FindTodo)
	t.DELETE("/:id", tc.DeleteTodo)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())

	e.Start(":8000")
}
