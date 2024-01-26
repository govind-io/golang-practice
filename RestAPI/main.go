package main

import (
	event "api/Models/Events"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/events", event.GetAllEventsFiber)

	app.Post("/events", event.SetEvent)

	app.Listen(":3000")

}
