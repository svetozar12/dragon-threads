package posts

import (
	"dragon-threads/apps/api/middleware"
	"dragon-threads/apps/api/middleware/fetch/subDragonFetch"

	"github.com/gofiber/fiber/v2"
)

func PostsRoute(app fiber.Router) {
	posts := app.Group("/posts")
	// Middlewares
	posts.Use(middleware.OAuth2Middleware, subDragonFetch.FetchSubDragon)

	// Routes
	posts.Get("/", getPostList)
	posts.Get("/:postId", GetPostById)
	posts.Post("/", createPost)
	posts.Put("/:postId", updatePost)
	posts.Delete("/:postId", deletePostById)
}
