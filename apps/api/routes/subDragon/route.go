package subDragon

import (
	"github.com/gofiber/fiber/v2"
)

func SubDragonsRoute(app fiber.Router) {
	SubDragons := app.Group("/subDragon")

	SubDragons.Get("/", getSubDragonList)
	SubDragons.Get("/:subDragonId", GetSubDragonById)
	SubDragons.Post("/", createSubDragon)
	SubDragons.Put("/:subDragonId", updateSubDragon)
	SubDragons.Delete("/:subDragonId", deleteSubDragonById)
}
