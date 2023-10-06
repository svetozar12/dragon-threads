package votes

import (
	"dragon-threads/apps/api/middleware"
	"dragon-threads/apps/api/middleware/fetch"

	"github.com/gofiber/fiber/v2"
)

func VotesRoute(app fiber.Router) {
	votes := app.Group("/posts")
	// Middlewares
	votes.Use(middleware.OAuth2Middleware, fetch.FetchSubDragon)

	// Routes
	votes.Post("/:postId/upvote", fetch.FetchSubDragon, fetch.FetchUser, votePost)
	votes.Post("/:postId/downvote", fetch.FetchSubDragon, fetch.FetchUser, votePost)

}
