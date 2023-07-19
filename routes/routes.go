package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/spknetwork/ecency_starting-point_be/handler"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
        return c.JSON(fiber.Map{
			"message": "Welcome to the break-away community hosting platform API",
		})
    })
	app.Post("/api/create-break", handler.CreateBreakaway)
}