package main

import (
	"first-task-alterra/configs"
	userHandler "first-task-alterra/internal/features/users/handler"
	userRepository "first-task-alterra/internal/features/users/repository"
	userServices "first-task-alterra/internal/features/users/services"

	todoHandler "first-task-alterra/internal/features/todos/handler"
	todoRepository "first-task-alterra/internal/features/todos/repository"
	todoServices "first-task-alterra/internal/features/todos/services"
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
	if err := db.AutoMigrate(&userRepository.User{}); err != nil {
		fmt.Println("Ada yg bermasalah saat memasukan table user", err.Error())
	}
	if err := db.AutoMigrate(&todoRepository.Todo{}); err != nil {
		fmt.Println("Ada yg bermasalah saat memasukan table todo", err.Error())
	}

	um := userRepository.NewUserModel(db)
	us := userServices.NewUserService(um)
	uc := userHandler.NewUserController(us)

	tm := todoRepository.NewTodoModel(db)
	ts := todoServices.NewTodoService(tm)
	tc := todoHandler.NewTodoController(ts)

	e := echo.New()
	e.GET("/hello", func(c echo.Context) error {
		return c.JSON(200, "hello world")
	})

	e.POST("/register", uc.Register())
	e.POST("/login", uc.Login())

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
	t.POST("", tc.AddTodo())
	t.PUT("/:id", tc.UpdateTodo())
	// t.GET("", tc.FindTodo)
	// t.DELETE("/:id", tc.DeleteTodo)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())

	e.Start(":8000")
}
