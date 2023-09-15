package users

import (
	"github.com/gofiber/fiber/v2"
)

func UsersRoute(app fiber.Router) {
	Users := app.Group("/users")
	Users.Get("/", getUserList)
	Users.Post("/", createUser)
}
