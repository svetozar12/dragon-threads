package auth

import (
	"github.com/gofiber/fiber/v2"
)

func AuthRoute(app fiber.Router) {
	Auth := app.Group("/auth")
	Auth.Get("/github", login)
	Auth.Get("/github/callback", githubCallback)
}
