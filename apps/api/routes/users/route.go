package users

import (
	"dragon-threads/apps/api/middleware"
	"dragon-threads/apps/api/middleware/fetch"

	"github.com/gofiber/fiber/v2"
)

func UsersRoute(app fiber.Router) {
	Users := app.Group("/users")
	// Middlewares
	Users.Use(middleware.OAuth2Middleware)

	// Routes
	Users.Get("/", getUserList)
	Users.Get("/:userId", fetch.FetchUser, GetUserById)
	Users.Put("/:userId", updateUser)
	Users.Delete("/:userId", deleteUserById)
}
