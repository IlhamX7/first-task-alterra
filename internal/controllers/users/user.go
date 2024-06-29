package users

import (
	"first-task-alterra/internal/helper"
	"first-task-alterra/internal/models"

	"first-task-alterra/internal/utils"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	model *models.UserModel
}

func NewUserController(m *models.UserModel) *UserController {
	return &UserController{
		model: m,
	}
}

func (uc *UserController) Login(c echo.Context) error {
	var input LoginRequest
	err := c.Bind(&input)
	if err != nil {
		return c.JSON(400, helper.ResponseFormat(400, "input error", nil))
	}
	result, err := uc.model.Login(input.Email, input.Password)
	if err != nil {
		return c.JSON(500, helper.ResponseFormat(500, "server error", nil))
	}
	token, err := utils.GenerateToken(result.ID)
	if err != nil {
		return c.JSON(500, helper.ResponseFormat(500, "privacy error", nil))
	}
	return c.JSON(200, helper.ResponseFormat(200, "success login", ToLoginReponse(result, token)))
}

func (uc *UserController) Register(c echo.Context) error {
	var input RegisterRequest
	err := c.Bind(&input)
	if err != nil {
		return c.JSON(400, helper.ResponseFormat(400, "input error", nil))
	}
	data, err := ToModelUsers(input)
	if err != nil {
		return c.JSON(400, helper.ResponseFormat(400, "input error", nil))
	}
	hashedPassword, err := utils.HashPassword(data.Password)
	if err != nil {
		return c.JSON(400, helper.ResponseFormat(400, "input error", nil))
	}
	_, err = uc.model.Register(data, hashedPassword)
	if err != nil {
		return c.JSON(500, helper.ResponseFormat(500, "server error", nil))
	}
	return c.JSON(201, helper.ResponseFormat(201, "success insert data", nil))
}
