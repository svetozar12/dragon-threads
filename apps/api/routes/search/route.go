package search

import (
	"dragon-threads/apps/api/middleware"

	"github.com/gofiber/fiber/v2"
)

func SearchRoute(app fiber.Router) {
	Search := app.Group("/search")
	// Middlewares
	Search.Use(middleware.OAuth2Middleware)

	// Routes
	Search.Get("/users", SearchUsers)

}
