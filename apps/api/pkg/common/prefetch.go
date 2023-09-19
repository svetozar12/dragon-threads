package common

import "github.com/gofiber/fiber/v2"

func Prefetch(resource string, value any, c *fiber.Ctx) {
	c.Locals(resource, value)
}

const SUB_DRAGON_BY_ID = "SUB_DRAGON_BY_ID"
