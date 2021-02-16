package routes

import (
	"github.com/gandhi-21/SampleFiberApi/controllers"
	"github.com/gofiber/fiber/v2"
)

func TodoRoute(route fiber.Router) {
	route.Get("", controllers.GetTodos)
}
