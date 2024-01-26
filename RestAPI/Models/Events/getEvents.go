package event

import "github.com/gofiber/fiber/v2"

func GetAllEventsFiber(c *fiber.Ctx) error {

	return c.JSON(GetAllEvents(), "application/json")
}
