package controllers

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// Todo ...
type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var todos = []*Todo{
	{
		ID:        1,
		Title:     "Walk the dog",
		Completed: false,
	},
	{
		ID:        2,
		Title:     "Walk the cat",
		Completed: false,
	},
}

// GetTodos ...
func GetTodos(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"todos": todos,
		},
	})
}

// GetTodo ...
func GetTodo(c *fiber.Ctx) error {
	paramID := c.Params("id")

	id, err := strconv.Atoi(paramID)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse ID",
		})
	}

	for _, todo := range todos {
		if todo.ID == id {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"success": true,
				"data": fiber.Map{
					"todo": todo,
				},
			})
		}
	}

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"success": false,
		"message": "Todo not found",
	})
}

// CreateTodo ...
func CreateTodo(c *fiber.Ctx) error {
	type Request struct {
		Title string `json:"title"`
	}

	var body Request

	err := c.BodyParser(&body)

	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
		})
	}

	todo := &Todo{
		ID:        len(todos) + 1,
		Title:     body.Title,
		Completed: false,
	}

	todos = append(todos, todo)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"todo": todo,
		},
	})
}
