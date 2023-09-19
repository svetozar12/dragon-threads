package subDragon

import (
	"dragon-threads/apps/api/middleware"

	"github.com/gofiber/fiber/v2"
)

func SubDragonsRoute(app fiber.Router) {
	SubDragons := app.Group("/subDragon")
	// Middlewares
	SubDragons.Use(middleware.OAuth2Middleware)
	// Routes
	SubDragons.Get("/", getSubDragonList)
	SubDragons.Get("/:subDragonId", GetSubDragonById)
	SubDragons.Post("/", createSubDragon)
	SubDragons.Put("/:subDragonId", updateSubDragon)
	SubDragons.Delete("/:subDragonId", deleteSubDragonById)
}
