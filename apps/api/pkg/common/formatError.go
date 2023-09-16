package common

import "github.com/gofiber/fiber/v2"

func FormatError(message string) fiber.Map {
	return fiber.Map{"error": message}
}
