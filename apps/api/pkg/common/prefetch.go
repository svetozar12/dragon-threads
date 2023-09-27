package common

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Prefetch(resource string, value any, c *fiber.Ctx) {
	c.Locals(resource, value)
}

func IsResourceFetched(resource string, c *fiber.Ctx) bool {
	res := c.Locals(resource)

	return res != nil
}

var SUB_DRAGON_BY_ID = func(id int) string {
	return fmt.Sprintf("SUB_DRAGON_BY_ID_%v", id)
}

var USER_BY_ID = func(id int) string {
	return fmt.Sprintf("USER_BY_ID_%v", id)
}
