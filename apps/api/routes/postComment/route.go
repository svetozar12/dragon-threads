package postComment

import (
	"dragon-threads/apps/api/middleware"
	"dragon-threads/apps/api/middleware/fetch"

	"github.com/gofiber/fiber/v2"
)

func PostCommentRoute(app fiber.Router) {
	postComment := app.Group("/posts/comment")
	// Middlewares
	postComment.Use(middleware.OAuth2Middleware, fetch.FetchSubDragon, fetch.FetchUser)

	// Routes
	postComment.Get("/", getPostList)
	postComment.Get("/:postId", GetPostById)
	postComment.Post("/", createPost)
	postComment.Put("/:postId", updatePost)
	postComment.Delete("/:postId", deletePostById)
}
