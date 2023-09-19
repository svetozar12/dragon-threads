package auth

import (
	"dragon-threads/apps/api/middleware/fetch/subDragonFetch"

	"github.com/gofiber/fiber/v2"
)

func AuthRoute(app fiber.Router) {
	Auth := app.Group("/auth")
	Auth.Get("/login", login)
	Auth.Get("/callback", subDragonFetch.FetchSubDragon, githubCallback)
}
