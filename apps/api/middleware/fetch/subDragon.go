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
	query := common.ParseIntQuery(c, "subDragonId", defaultValue)
	param := common.ParseIntParam(c, "subDragonId", defaultValue)
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

	// Check if the "userId" is provided in the request body
	userIdFromBody, ok := requestBody["postId"].(int)
	if !ok {
		userIdFromBody = 0 // If userId is not provided or not an integer, use a default value
	}
	return userIdFromBody, nil
}
