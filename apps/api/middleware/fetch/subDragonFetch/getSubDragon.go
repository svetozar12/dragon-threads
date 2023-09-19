package subDragonFetch

import (
	"dragon-threads/apps/api/constants"
	"dragon-threads/apps/api/pkg/common"
	"dragon-threads/apps/api/repositories/subDragonRepository"

	"github.com/gofiber/fiber/v2"
)

func FetchSubDragon(c *fiber.Ctx) error {
	defaultValue := 0
	query := common.ParseIntQuery(c, "subDragonId", defaultValue)
	param := common.ParseIntParam(c, "subDragonId", defaultValue)
	subDragonId := query | param
	value, err := subDragonRepository.GetSubDragon("id=?", subDragonId)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(common.FormatError(constants.SUB_DRAGON_NOT_FOUND))
	}
	common.Prefetch(common.SUB_DRAGON_BY_ID, value, c)
	// Call the next middleware or handler
	return c.Next()
}
