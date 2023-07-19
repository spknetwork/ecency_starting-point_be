package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/spknetwork/ecency_starting-point_be/routes"
)

func main() {
    app := fiber.New()

    
    routes.SetupRoutes(app)
    app.Listen(":3000")
}