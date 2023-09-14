package users

import "github.com/gofiber/fiber/v2"

const USER_ALREADY_EXIST = "User already exist"
const USER_FAILED_TO_CREATE = "Failed to create user"

func formatError(message string) fiber.Map {
	return fiber.Map{"error": message}
}
