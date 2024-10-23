package services

import (
	"errors"

	"github.com/bancodebogota/bbog-dig-pl-go-mngr-template/src/app/dal"
	"github.com/bancodebogota/bbog-dig-pl-go-mngr-template/src/app/models"
	"github.com/bancodebogota/bbog-dig-pl-go-mngr-template/src/app/types"
	"github.com/bancodebogota/bbog-dig-pl-go-mngr-template/src/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateTodo(c *fiber.Ctx) error {
	b := new(types.CreateDTO)

	if err := utils.ParseBodyAndValidate(c, b); err != nil {
		return err
	}

	d := &models.Todo{
		Task: b.Task,
		User: utils.GetUser(c),
	}

	if err := dal.CreateTodo(d).Error; err != nil {
		return fiber.NewError(fiber.StatusConflict, err.Error())
	}

	return c.JSON(&types.TodoCreateResponse{
		Todo: &types.TodoResponse{
			ID:        d.ID,
			Task:      d.Task,
			Completed: d.Completed,
		},
	})
}

func GetTodos(c *fiber.Ctx) error {
	d := &[]types.TodoResponse{}

	err := dal.FindTodosByUser(d, utils.GetUser(c)).Error
	if err != nil {
		return fiber.NewError(fiber.StatusConflict, err.Error())
	}

	return c.JSON(&types.TodosResponse{
		Todos: d,
	})
}

func GetTodo(c *fiber.Ctx) error {
	todoID := c.Params("todoID")

	if todoID == "" {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "Invalid todoID")
	}

	d := &types.TodoResponse{}

	err := dal.FindTodoByUser(d, todoID, utils.GetUser(c)).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.JSON(&types.TodoCreateResponse{})
	}

	return c.JSON(&types.TodoCreateResponse{
		Todo: d,
	})
}

func DeleteTodo(c *fiber.Ctx) error {
	todoID := c.Params("todoID")

	if todoID == "" {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "Invalid todoID")
	}

	res := dal.DeleteTodo(todoID, utils.GetUser(c))
	if res.RowsAffected == 0 {
		return fiber.NewError(fiber.StatusConflict, "Unable to delete todo")
	}

	err := res.Error
	if err != nil {
		return fiber.NewError(fiber.StatusConflict, err.Error())
	}

	return c.JSON(&types.MsgResponse{
		Message: "Todo successfully deleted",
	})
}

func CheckTodo(c *fiber.Ctx) error {
	b := new(types.CheckTodoDTO)
	todoID := c.Params("todoID")

	if todoID == "" {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "Invalid todoID")
	}

	if err := utils.ParseBodyAndValidate(c, b); err != nil {
		return err
	}

	err := dal.UpdateTodo(todoID, utils.GetUser(c), map[string]interface{}{"completed": b.Completed}).Error
	if err != nil {
		return fiber.NewError(fiber.StatusConflict, err.Error())
	}

	return c.JSON(&types.MsgResponse{
		Message: "Todo successfully updated",
	})
}

func UpdateTodoTitle(c *fiber.Ctx) error {
	b := new(types.CreateDTO)
	todoID := c.Params("todoID")

	if todoID == "" {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "Invalid todoID")
	}

	if err := utils.ParseBodyAndValidate(c, b); err != nil {
		return err
	}

	err := dal.UpdateTodo(todoID, utils.GetUser(c), &models.Todo{Task: b.Task}).Error
	if err != nil {
		return fiber.NewError(fiber.StatusConflict, err.Error())
	}

	return c.JSON(&types.MsgResponse{
		Message: "Todo successfully updated",
	})
}
