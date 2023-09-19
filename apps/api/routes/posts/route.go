package posts

import (
	"dragon-threads/apps/api/middleware"

	"github.com/gofiber/fiber/v2"
)

func PostsRoute(app fiber.Router) {
	Posts := app.Group("/posts")
	// Middlewares
	Posts.Use(middleware.OAuth2Middleware)
	// Routes
	Posts.Get("/", getPostList)
	Posts.Get("/:postId", GetPostById)
	Posts.Post("/", createPost)
	Posts.Put("/:postId", updatePost)
	Posts.Delete("/:postId", deletePostById)
}
