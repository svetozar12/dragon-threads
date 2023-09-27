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
	posts.Get("/", fetch.FetchSubDragon, getPostList)
	posts.Get("/:postId", fetch.FetchSubDragon, GetPostById)
	posts.Post("/", fetch.FetchSubDragon, fetch.FetchUser, createPost)
	posts.Put("/:postId", fetch.FetchSubDragon, fetch.FetchUser, updatePost)
	posts.Delete("/:postId", fetch.FetchSubDragon, deletePostById)
}
