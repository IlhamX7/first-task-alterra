package factory

import (
	"first-task-alterra/configs"
	userHandler "first-task-alterra/internal/features/users/handler"
	userRepository "first-task-alterra/internal/features/users/repository"
	userServices "first-task-alterra/internal/features/users/services"

	todoHandler "first-task-alterra/internal/features/todos/handler"
	todoRepository "first-task-alterra/internal/features/todos/repository"
	todoServices "first-task-alterra/internal/features/todos/services"
	"first-task-alterra/internal/routes"
	"fmt"

	"first-task-alterra/internal/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitFactory(e *echo.Echo) {
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

	pu := utils.NewPasswordUtility()
	um := userRepository.NewUserModel(db)
	us := userServices.NewUserService(um, pu)
	uc := userHandler.NewUserController(us)

	tm := todoRepository.NewTodoModel(db)
	ts := todoServices.NewTodoService(tm)
	tc := todoHandler.NewTodoController(ts)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())

	routes.InitRoute(e, uc, tc)
}
