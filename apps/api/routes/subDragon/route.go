package subDragon

import (
	"dragon-threads/apps/api/middleware"
	"dragon-threads/apps/api/middleware/fetch/subDragonFetch"

	"github.com/gofiber/fiber/v2"
)

func SubDragonsRoute(app fiber.Router) {
	subDragons := app.Group("/subDragon")
	// Middlewares
	subDragons.Use(middleware.OAuth2Middleware)
	// Routes
	subDragons.Get("/", getSubDragonList)
	subDragons.Get("/:subDragonId", subDragonFetch.FetchSubDragon, getSubDragonById)
	subDragons.Post("/", createSubDragon)
	subDragons.Put("/:subDragonId", subDragonFetch.FetchSubDragon, updateSubDragon)
	subDragons.Delete("/:subDragonId", subDragonFetch.FetchSubDragon, deleteSubDragonById)
}
