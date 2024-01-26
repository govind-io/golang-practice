package event

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func SetEvent(c *fiber.Ctx) error {
	newEvent := Event{}
	err := c.BodyParser(&newEvent)

	if err != nil {
		fmt.Println(err)
		return c.Status(400).JSON(map[string]string{"message": "Invalid req body"})
	}

	newEvent.Save()

	return c.JSON(map[string][]Event{"data": events})
}
