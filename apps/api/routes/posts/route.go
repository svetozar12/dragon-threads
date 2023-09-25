package posts

import (
	"dragon-threads/apps/api/middleware"
	"dragon-threads/apps/api/middleware/fetch"

	"github.com/gofiber/fiber/v2"
)

func PostsRoute(app fiber.Router) {
	posts := app.Group("/posts")
	// Middlewares
	posts.Use(middleware.OAuth2Middleware, fetch.FetchSubDragon)

	// Routes
	posts.Get("/", getPostList)
	posts.Get("/:postId", GetPostById)
	posts.Post("/", fetch.FetchUser, createPost)
	posts.Put("/:postId", fetch.FetchUser, updatePost)
	posts.Delete("/:postId", deletePostById)
}
