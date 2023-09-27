package fetch

import (
	"dragon-threads/apps/api/constants"
	"dragon-threads/apps/api/pkg/common"
	"dragon-threads/apps/api/repositories/usersRepository"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func FetchUser(c *fiber.Ctx) error {
	defaultValue := 0
	query := common.ParseIntQuery(c, constants.USER_ID, defaultValue)
	param := common.ParseIntParam(c, constants.USER_ID, defaultValue)
	body, _ := getUserIdFromBody(c)
	userId := query | param | body
	if isFetched := common.IsResourceFetched(common.USER_BY_ID(userId), c); isFetched {
		return c.Next()
	}
	value, err := usersRepository.GetUser("id=?", userId)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(common.FormatError(constants.USER_NOT_FOUND))
	}
	common.Prefetch(common.USER_BY_ID(userId), value, c)
	// Call the next middleware or handler
	return c.Next()
}

func getUserIdFromBody(c *fiber.Ctx) (int, error) {
	var requestBody map[string]interface{}
	if err := c.BodyParser(&requestBody); err != nil {
		return 0, fmt.Errorf("Invalid request body")
	}

	// Check if the "userId" is provided in the request body
	userIdFromBody, ok := requestBody["userId"].(int)
	if !ok {
		userIdFromBody = 0 // If userId is not provided or not an integer, use a default value
	}
	return userIdFromBody, nil
}
