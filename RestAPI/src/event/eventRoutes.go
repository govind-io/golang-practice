package event

import (
	"github.com/gofiber/fiber/v2"
)

func GetEventApp() *fiber.App {
	EventApp := fiber.New()

	EventApp.Get("/", GetAllEventsFiber)

	EventApp.Post("/", SetEvent)

	return EventApp
}
