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
	posts.Get("/", fetch.FetchSubDragon, fetch.FetchUser, getPostList)
	posts.Get("/:postId", fetch.FetchSubDragon, fetch.FetchUser, GetPostById)
	posts.Post("/", fetch.FetchSubDragon, fetch.FetchUser, createPost)
	posts.Put("/:postId", fetch.FetchSubDragon, fetch.FetchUser, updatePost)
	posts.Delete("/:postId", fetch.FetchSubDragon, fetch.FetchUser, deletePostById)
}
