package postComment

import (
	"dragon-threads/apps/api/middleware"
	"dragon-threads/apps/api/middleware/fetch"

	"github.com/gofiber/fiber/v2"
)

func PostCommentRoute(app fiber.Router) {
	postComment := app.Group("/posts/comment")
	// Middlewares
	postComment.Use(middleware.OAuth2Middleware, fetch.FetchSubDragon, fetch.FetchUser, fetch.FetchPost)

	// Routes
	postComment.Get("/", getPostCommentList)
	postComment.Get("/:postId", GetPostCommentById)
	postComment.Post("/", createPostComment)
	postComment.Put("/:postId", updatePostComment)
	postComment.Delete("/:postId", deletePostCommentById)
}
