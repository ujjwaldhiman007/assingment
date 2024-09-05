package routes

import (
	"gofiber-app/internal/app/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/api/todos", handler.GetCars)
	app.Post("/api/todos", handler.CreateCar)
	app.Patch("/api/todos/:id", handler.UpdateCar)
	app.Delete("/api/todos/:id", handler.DeleteCar)
}
