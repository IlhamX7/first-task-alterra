package todos

import (
	"first-task-alterra/internal/helper"
	"first-task-alterra/internal/models"

	"first-task-alterra/internal/utils"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type TodoController struct {
	model *models.TodoModel
}

func NewTodoController(m *models.TodoModel) *TodoController {
	return &TodoController{
		model: m,
	}
}

func (tc *TodoController) AddTodo(c echo.Context) error {
	var input CreateRequest
	err := c.Bind(&input)
	if err != nil {
		return c.JSON(400, helper.ResponseFormat(400, "input error", nil))
	}
	var userID = utils.DecodeToken(c.Get("user").(*jwt.Token))
	_, err = tc.model.AddTodo(input.Activity, userID)
	if err != nil {
		return c.JSON(500, helper.ResponseFormat(500, "server error", nil))
	}
	return c.JSON(201, helper.ResponseFormat(201, "success insert data", nil))
}

func (tc *TodoController) UpdateTodo(c echo.Context) error {
	id := c.Param("id")
	getId, err := helper.StringToUint(id)
	if err != nil {
		return c.JSON(400, helper.ResponseFormat(400, "input error", nil))
	}
	var input UpdateRequest

	err = c.Bind(&input)
	if err != nil {
		return c.JSON(400, helper.ResponseFormat(400, "input error", nil))
	}
	_, err = tc.model.UpdateTodo(getId, input.Activity, input.Mark, input.UserId)
	if err != nil {
		return c.JSON(500, helper.ResponseFormat(500, "server error", nil))
	}
	return c.JSON(200, helper.ResponseFormat(200, "success update data", nil))
}

func (tc *TodoController) DeleteTodo(c echo.Context) error {
	id := c.Param("id")
	getId, err := helper.StringToUint(id)
	if err != nil {
		return c.JSON(400, helper.ResponseFormat(400, "input error", nil))
	}
	_, err = tc.model.DeleteTodo(getId)
	if err != nil {
		return c.JSON(500, helper.ResponseFormat(500, "server error", nil))
	}
	return c.JSON(200, helper.ResponseFormat(200, "success delete data", nil))
}

func (tc *TodoController) FindTodo(c echo.Context) error {
	var userID = utils.DecodeToken(c.Get("user").(*jwt.Token))
	data, err := tc.model.FindTodo(userID)
	if err != nil {
		return c.JSON(500, helper.ResponseFormat(500, "server error", nil))
	}
	return c.JSON(200, helper.ResponseFormat(200, "success get all data", data))
}
