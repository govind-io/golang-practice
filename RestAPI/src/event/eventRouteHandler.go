package event

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetAllEventsFiber(c *fiber.Ctx) error {

	return c.JSON(GetAllEvents(), "application/json")
}

func SetEvent(c *fiber.Ctx) error {
	newEvent := Event{}
	err := c.BodyParser(&newEvent)

	if err != nil {
		return c.Status(400).JSON(map[string]string{"message": "Invalid req body"})
	}

	newEvent.DateTime = time.Now()

	newEvent.Save()

	return c.JSON(map[string][]Event{"data": events})
}
