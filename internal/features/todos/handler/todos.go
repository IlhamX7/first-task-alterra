package handler

import (
	"first-task-alterra/internal/features/todos"
	"first-task-alterra/internal/helper"
	"first-task-alterra/internal/utils"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type TodoController struct {
	srv todos.Services
}

func NewTodoController(s todos.Services) todos.Handler {
	return &TodoController{
		srv: s,
	}
}

func (tc *TodoController) AddTodo() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input CreateRequest
		err := c.Bind(&input)
		if err != nil {
			return c.JSON(400, helper.ResponseFormat(400, "input error", nil))
		}
		var userID = utils.DecodeToken(c.Get("user").(*jwt.Token))

		err = tc.srv.AddTodo(input.Activity, userID)
		if err != nil {
			return c.JSON(500, helper.ResponseFormat(500, "server error", nil))
		}
		return c.JSON(201, helper.ResponseFormat(201, "success insert data", nil))
	}
}

func (tc *TodoController) UpdateTodo() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		getId, err := utils.StringToUint(id)
		if err != nil {
			return c.JSON(400, helper.ResponseFormat(400, "input error", nil))
		}
		var input UpdateRequest

		err = c.Bind(&input)
		if err != nil {
			return c.JSON(400, helper.ResponseFormat(400, "input error", nil))
		}
		err = tc.srv.UpdateTodo(getId, input.Activity, input.Mark, input.UserId)
		if err != nil {
			return c.JSON(500, helper.ResponseFormat(500, "server error", nil))
		}
		return c.JSON(200, helper.ResponseFormat(200, "success update data", nil))
	}
}
