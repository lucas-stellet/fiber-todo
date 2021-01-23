package controllers

import "github.com/gofiber/fiber/v2"

// Todo ...
type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var todos = []Todo{
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
