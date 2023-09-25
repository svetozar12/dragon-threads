package subDragon

import (
	"dragon-threads/apps/api/middleware"
	"dragon-threads/apps/api/middleware/fetch"

	"github.com/gofiber/fiber/v2"
)

func SubDragonsRoute(app fiber.Router) {
	subDragons := app.Group("/subDragon")
	// Middlewares
	subDragons.Use(middleware.OAuth2Middleware)
	// Routes
	subDragons.Get("/", getSubDragonList)
	subDragons.Get("/:subDragonId", fetch.FetchSubDragon, getSubDragonById)
	subDragons.Post("/", createSubDragon)
	subDragons.Put("/:subDragonId", fetch.FetchSubDragon, updateSubDragon)
	subDragons.Delete("/:subDragonId", fetch.FetchSubDragon, deleteSubDragonById)
}
