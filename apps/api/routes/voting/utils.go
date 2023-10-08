package votes

import (
	"dragon-threads/apps/api/entities"
	"dragon-threads/apps/api/pkg/common"
	"dragon-threads/apps/api/repositories/postVoteRepository"

	"github.com/gofiber/fiber/v2"
)

func handleVote(c *fiber.Ctx, prePost *entities.Post, preUser *entities.User) error {
	// Determine the vote type based on the endpoint (upvote or downvote)
	voteValue := 1
	if c.Route().Path == "/v1/posts/{id}/downvote" {
		voteValue = -1
	}
	postVote, err := postVoteRepository.GetPostVote("user_id =? AND post_id =?", preUser.ID, prePost.ID)
	if err != nil {
		_, err := postVoteRepository.CreatePostVote(&entities.PostVote{UserID: int32(preUser.ID), PostID: int32(prePost.ID), Vote: int32(voteValue)})

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(common.FormatError(err.Error()))
		}
	}
	if voteValue == int(postVote.Vote) {
		return c.Status(fiber.StatusBadRequest).JSON(common.FormatError("You have already upvoted/downvoted this post"))
	}
	postVote.Vote += int32(voteValue)
	return nil
}
