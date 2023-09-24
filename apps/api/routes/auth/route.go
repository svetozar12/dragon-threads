package auth

import (
	"github.com/gofiber/fiber/v2"
)

func AuthRoute(app fiber.Router) {
	Auth := app.Group("/auth")
	Auth.Get("/login", login)
	Auth.Get("/callback", githubCallback)
}
