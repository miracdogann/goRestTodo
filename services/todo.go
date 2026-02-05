package services

import (
	"errors"
	"fiber_rest/dal"
	"fiber_rest/types"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

var validate = validator.New()

func GetTodos(c fiber.Ctx) error {
	todos := []types.TodoResponse{}
	res := dal.GetTodos(&todos)

	if res.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to get todos",
		})
	}
	return c.JSON(todos)
}

func GetTodo(c fiber.Ctx) error {

	todoID := c.Params("todoID")

	d := types.TodoResponse{}
	res := dal.GetTodoByID(&d, todoID)

	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return c.Status(404).JSON(fiber.Map{
				"message": "Todo not found",
			})
		}
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to get todo",
		})
	}
	return c.JSON(d)
}

func CreateTodo(c fiber.Ctx) error {

	t := new(types.TodoCreateDTO)

	if err := c.Bind().Body(t); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}

	if err := validate.Struct(t); err != nil {
		valErr := err.(validator.ValidationErrors)[0]
		message := fmt.Sprintf("Field : '%s' , failed on '%s' with your value : '%s'", valErr.Field(), valErr.Tag(), valErr.Value())

		return c.Status(400).JSON(fiber.Map{
			"message": message,
		})
	}

	newTodo := dal.Todo{
		Title: t.Title,
	}
	res := dal.CreateTodo(&newTodo)
	if res.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to create to do",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Todo kaydı başarılı",
	})
}
func UpdateTodo(c fiber.Ctx) error {
	todoID := c.Params("todoID")

	t := new(types.TodoUpdateDTO)

	if err := c.Bind().Body(t); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Bad Reequst",
		})
	}
	if err := validate.Struct(t); err != nil {
		valErr := err.(validator.ValidationErrors)[0]
		message := fmt.Sprintf("Field : '%s' , failed on '%s' with your value : '%s'", valErr.Field(), valErr.Tag(), valErr.Value())

		return c.Status(400).JSON(fiber.Map{
			"message": message,
		})
	}
	res := dal.UpdateTodo(todoID, t)
	if res.Error != nil || res.RowsAffected == 0 {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to updated todo",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Todo Updated Successfully",
	})
}

func DeleteTodo(c fiber.Ctx) error {
	todoID := c.Params("todoID")
	res := dal.DeleteTodo(todoID)

	if res.Error != nil || res.RowsAffected == 0 {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to delete todo",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Todo Deleted Sucessfully",
	})
}
