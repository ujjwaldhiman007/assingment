package routes

import (
	"gofiber-app/internal/app/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/api/cars", handler.GetCars)
	app.Post("/api/cars", handler.CreateCar)
	app.Patch("/api/cars/:id", handler.UpdateCar)
	app.Delete("/api/cars/:id", handler.DeleteCar)
}
