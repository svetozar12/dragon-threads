package votes

import (
	"dragon-threads/apps/api/entities"
	"dragon-threads/apps/api/pkg/common"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// Refactored godoc
// @Summary      Vote on a Post
// @Tags         Post
// @Accept       json
// @Param request body votes.PostVoteSchema true "vote request"
// @Security ApiKeyAuth
// @Success      201  {object} entities.Post
// @Failure      400  {object} common.CommonErrorSchema
// @Router       /v1/posts/{id}/upvote [post]
func upVotePost(c *fiber.Ctx) error {

	var postVote PostVoteSchema
	if err := c.BodyParser(&postVote); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.FormatError(err.Error()))
	}
	prePost := c.Locals(common.POST_BY_ID(int(postVote.PostID))).(*entities.Post)
	preUser := c.Locals(common.POST_BY_ID(int(postVote.UserID))).(*entities.User)

	validate := validator.New()
	if err := validate.Struct(postVote); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.FormatError(err.Error()))
	}

	err := handleVote(c, prePost, preUser)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(postVote)
}

// Refactored godoc
// @Summary      Vote on a Post
// @Tags         Post
// @Accept       json
// @Param request body votes.PostVoteSchema true "vote request"
// @Security ApiKeyAuth
// @Success      201  {object} entities.Post
// @Failure      400  {object} common.CommonErrorSchema
// @Router       /v1/posts/{id}/downvote [post]
func downVotePost(c *fiber.Ctx) error {
	var postVote PostVoteSchema
	if err := c.BodyParser(&postVote); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.FormatError(err.Error()))
	}
	prePost := c.Locals(common.POST_BY_ID(int(postVote.PostID))).(*entities.Post)
	preUser := c.Locals(common.POST_BY_ID(int(postVote.UserID))).(*entities.User)

	validate := validator.New()
	if err := validate.Struct(postVote); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.FormatError(err.Error()))
	}

	err := handleVote(c, prePost, preUser)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(postVote)
}
