package routes

import (
	"dragon-threads/apps/api/internal/routes/users"

	_ "dragon-threads/apps/api/internal/pkg/swagger"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3333
// @BasePath /
func InitRoutes(app *fiber.App) {
	v1 := app.Group("/v1")
	v1.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello from the api")
	})
	users.UsersRoute(v1)
	v1.Get("/swagger/*", swagger.HandlerDefault) // default
}
