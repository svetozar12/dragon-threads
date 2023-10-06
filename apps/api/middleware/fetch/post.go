package fetch

import (
	"dragon-threads/apps/api/constants"
	"dragon-threads/apps/api/pkg/common"
	"dragon-threads/apps/api/repositories/postsRepository"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func FetchPost(c *fiber.Ctx) error {
	defaultValue := 0
	query := common.ParseIntQuery(c, constants.POST_ID, defaultValue)
	param := common.ParseIntParam(c, constants.POST_ID, defaultValue)
	body, _ := getSubdragonIdFromBody(c)

	postId := query | param | body
	if isFetched := common.IsResourceFetched(common.POST_BY_ID(postId), c); isFetched {
		return c.Next()
	}
	value, err := postsRepository.GetPost("id=?", postId)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(common.FormatError(constants.SUB_DRAGON_NOT_FOUND))
	}
	common.Prefetch(common.SUB_DRAGON_BY_ID(postId), value, c)
	// Call the next middleware or handler
	return c.Next()
}

func getPostIdFromBody(c *fiber.Ctx) (int, error) {
	var requestBody map[string]interface{}
	if err := c.BodyParser(&requestBody); err != nil {
		return 0, fmt.Errorf("Invalid request body")
	}

	// Check if the "postId" is provided in the request body
	postIdFromBody, ok := requestBody[constants.POST_ID].(int)
	if !ok {
		postIdFromBody = 0 // If postId is not provided or not an integer, use a default value
	}
	return postIdFromBody, nil
}
