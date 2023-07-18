package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/spknetwork/ecency_starting-point_be/handler"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/api/create-break", handler.CreateBreakaway)
}