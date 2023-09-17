package posts

import (
	"github.com/gofiber/fiber/v2"
)

func PostsRoute(app fiber.Router) {
	Posts := app.Group("/posts")

	Posts.Get("/", getPostList)
	Posts.Get("/:postId", GetPostById)
	Posts.Post("/", createPost)
	Posts.Put("/:postId", updatePost)
	Posts.Delete("/:postId", deletePostById)
}
