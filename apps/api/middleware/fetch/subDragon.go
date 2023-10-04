package fetch

import (
	"dragon-threads/apps/api/constants"
	"dragon-threads/apps/api/pkg/common"
	"dragon-threads/apps/api/repositories/subDragonRepository"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func FetchSubDragon(c *fiber.Ctx) error {
	defaultValue := 0
	query := common.ParseIntQuery(c, constants.SUB_DRAGON_ID, defaultValue)
	param := common.ParseIntParam(c, constants.SUB_DRAGON_ID, defaultValue)
	body, _ := getSubdragonIdFromBody(c)

	subDragonId := query | param | body
	if isFetched := common.IsResourceFetched(common.SUB_DRAGON_BY_ID(subDragonId), c); isFetched {
		return c.Next()
	}
	value, err := subDragonRepository.GetSubDragon("id=?", subDragonId)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(common.FormatError(constants.SUB_DRAGON_NOT_FOUND))
	}
	common.Prefetch(common.SUB_DRAGON_BY_ID(subDragonId), value, c)
	// Call the next middleware or handler
	return c.Next()
}

func getSubdragonIdFromBody(c *fiber.Ctx) (int, error) {
	var requestBody map[string]interface{}
	if err := c.BodyParser(&requestBody); err != nil {
		return 0, fmt.Errorf("Invalid request body")
	}

	// Check if the "subDragonId" is provided in the request body
	subDragonIdFromBody, ok := requestBody[constants.SUB_DRAGON_ID].(int)
	if !ok {
		subDragonIdFromBody = 0 // If subDragonId is not provided or not an integer, use a default value
	}
	return subDragonIdFromBody, nil
}
