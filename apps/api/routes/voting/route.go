package votes

import (
	"dragon-threads/apps/api/middleware"
	"dragon-threads/apps/api/middleware/fetch"

	"github.com/gofiber/fiber/v2"
)

func VotesRoute(app fiber.Router) {
	votes := app.Group("/posts")
	// Middlewares
	votes.Use(middleware.OAuth2Middleware, fetch.FetchSubDragon, fetch.FetchPost, fetch.FetchUser)

	// Routes
	votes.Post("/:postId/upvote", upVotePost)
	votes.Post("/:postId/downvote", downVotePost)

}
