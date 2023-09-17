package posts

import (
	"dragon-threads/apps/api/entities"
	"dragon-threads/apps/api/repositories/postsRepository"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func parsePaginationQuery(c *fiber.Ctx) (int, int) {
	pageStr := c.Query("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}

	pageSizeStr := c.Query("pageSize")
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		pageSize = 10
	}

	return page, pageSize
}

func parseGetByQuery(c *fiber.Ctx) (string, string) {
	getBy := c.Query("getBy")
	getByValue := c.Query("getByValue")
	if getBy == "" || getByValue == "" {
		getBy = "1"
		getByValue = "1"
	}

	return getBy, getByValue
}

func parsePostIdParams(c *fiber.Ctx) (int, error) {
	userIdStr := c.Params("userId")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		return 0, fmt.Errorf("Post id is required")
	}
	return userId, nil
}
func getPostsByQuery(getBy string, getByValue string, page int, pageSize int) ([]entities.Post, int64, error) {
	// Retrieve user list based on query parameters
	userList, total, err := postsRepository.GetPostList(getBy+"=?", page, pageSize, []interface{}{getByValue})
	return userList, total, err
}

func updatePostFields(existingPost *entities.Post, updatedPost UpdatePostSchema) {
	if updatedPost.Title != "" {
		existingPost.Title = updatedPost.Title
	}

	if updatedPost.Content != "" {
		existingPost.Content = updatedPost.Content
	}
}
