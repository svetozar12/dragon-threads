package auth

import (
	"dragon-threads/apps/api/middleware"

	"github.com/gofiber/fiber/v2"
)

func AuthRoute(app fiber.Router) {
	Auth := app.Group("/auth")
	Auth.Get("/github", login)
	Auth.Get("/github/callback", githubCallback)
	Auth.Get("/verify", middleware.OAuth2Middleware, verifyToken)
}
