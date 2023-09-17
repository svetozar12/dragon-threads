package subDragon

import (
	"github.com/gofiber/fiber/v2"
)

func SubDragonsRoute(app fiber.Router) {
	SubDragons := app.Group("/users")

	SubDragons.Get("/", getSubDragonList)
	SubDragons.Get("/:userId", GetSubDragonById)
	SubDragons.Post("/", createSubDragon)
	SubDragons.Put("/:userId", updateSubDragon)
	SubDragons.Delete("/:userId", deleteSubDragonById)
}
