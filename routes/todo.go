package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lucas-stellet/fiber-todo/controllers"
)

// TodoRoute ...
func TodoRoute(route fiber.Router) {
	route.Get("", controllers.GetTodos)
}
