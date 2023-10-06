package votes

import (
	"dragon-threads/apps/api/constants"
	"dragon-threads/apps/api/entities"
	"dragon-threads/apps/api/pkg/common"
	"dragon-threads/apps/api/repositories/postsRepository"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// Refactored godoc
// @Summary      Vote on a Post
// @Tags         Post
// @Accept       json
// @Param request body posts.PostVoteSchema true "vote request"
// @Security ApiKeyAuth
// @Success      201  {object} entities.Post
// @Failure      400  {object} common.CommonErrorSchema
// @Router       /v1/posts/{id}/vote [post]
func votePost(c *fiber.Ctx) error {
	postID := common.ParseIntParam(c, constants.POST_ID, 0)
	prePost := c.Locals(common.POST_BY_ID(postID)).(*entities.Post)

	var vote PostVoteSchema
	if err := c.BodyParser(&vote); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.FormatError(err.Error()))
	}

	validate := validator.New()
	if err := validate.Struct(vote); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.FormatError(err.Error()))
	}

	// Determine the vote type based on the endpoint (upvote or downvote)
	voteValue := 1
	if c.Route().Path == "/v1/posts/{id}/downvote" {
		voteValue = -1
	}

	prePost.Votes += int32(voteValue)
	_, err := postsRepository.UpdatePost(prePost)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.FormatError(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(vote)
}
