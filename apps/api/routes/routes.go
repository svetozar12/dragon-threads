package routes

import (
	"dragon-threads/apps/api/routes/auth"
	"dragon-threads/apps/api/routes/posts"
	"dragon-threads/apps/api/routes/search"
	"dragon-threads/apps/api/routes/subDragon"
	"dragon-threads/apps/api/routes/users"
	"fmt"

	_ "dragon-threads/apps/api/pkg/swagger"

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
// @securityDefinitions.apikey ApiKeyAuth
// @In header
// @Name Authorization
// @BasePath /
func InitRoutes(app *fiber.App) {
	v1 := app.Group("/v1")
	v1.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello from the api")
	})
	users.UsersRoute(v1)
	posts.PostsRoute(v1)
	subDragon.SubDragonsRoute(v1)
	auth.AuthRoute(v1)
	search.SearchRoute(v1)
	v1.Get("/swagger/*", swagger.HandlerDefault) // default
	fmt.Println("Routes Initialized")
}
