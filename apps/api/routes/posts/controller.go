package posts

import (
	"dragon-threads/apps/api/entities"
	"dragon-threads/apps/api/pkg/common"
	"dragon-threads/apps/api/repositories/postsRepository"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// @Summary Get post by ID
// @Tags Post
// @Accept json
// @Param postId path int false "Post ID"
// @Success 200 {object} posts.PostSchema "Success"
// @Failure 400 {object} common.CommonErrorSchema "Bad Request"
// @Failure 404 {object} common.CommonErrorSchema "Not Found Request"
// @Router /v1/posts/{postId} [get]
func GetPostById(c *fiber.Ctx) error {
	// Parse pagination parameters
	postID, err := parsePostIdParams(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.FormatError(err.Error()))
	}
	// Get post list based on query parameters
	post, err := postsRepository.GetPost("id =?", postID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(common.FormatError(POST_NOT_FOUND))
	}
	// Construct the response
	response := post

	return c.Status(fiber.StatusOK).JSON(response)
}

// @Summary      Get Post List
// @Tags         Post
// @Accept       json
// @Param page       query int false "Page number (default: 1)"
// @Param pageSize   query int false "Number of items per page (default: 10)"
// @Param getBy      query GetByEnum false "Get posts by field (optional)"
// @Param getByValue query int false "Get posts by field value (optional)".
// @Success      200  {object} posts.PostListSchema "Success"
// @Failure 	 400  {object} common.CommonErrorSchema
// @Router       /v1/posts [get]
func getPostList(c *fiber.Ctx) error {
	// Parse pagination parameters
	page, pageSize := parsePaginationQuery(c)

	// Get query and value parameters
	getBy, getByValue := parseGetByQuery(c)

	// Get post list based on query parameters
	postList, total, err := getPostsByQuery(getBy, getByValue, page, pageSize)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.FormatError(err.Error()))
	}

	// Check if there is a next page
	hasNextPage := (page * pageSize) < int(total)

	// Construct the response
	response := PostListSchema{
		Data: postList,
		Pagination: common.CommonPaginationSchema{
			Page:       page,
			PageSize:   pageSize,
			TotalCount: total,
			HasNext:    hasNextPage,
		},
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

// Example godoc
// @Summary      Create Post
// @Tags         Post
// @Accept       json
// @Param request body posts.PostSchema true "query params""
// @Success      201  {object} entities.Post
// @Failure 	 400  {object} common.CommonErrorSchema
// @Router       /v1/posts [post]
func createPost(c *fiber.Ctx) error {
	var post PostSchema
	if err := c.BodyParser(&post); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.FormatError(err.Error()))
	}
	validate := validator.New()
	if err := validate.Struct(post); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.FormatError(err.Error()))
	}

	_, err := postsRepository.CreatePost(&entities.Post{
		Title:       post.Title,
		Content:     post.Content,
		UserID:      post.UserID,
		SubDragonId: post.SubDragonId,
	})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.FormatError(err.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(post)
}

// @Summary      Update Post
// @Tags         Post
// @Accept       json
// @Param id path string true "Post ID"
// @Param request body posts.UpdatePostSchema true "Request body for updating post"
// @Success      200  {object} entities.Post
// @Failure      400  {object} common.CommonErrorSchema
// @Router       /v1/posts/{id} [put]
func updatePost(c *fiber.Ctx) error {
	// Retrieve the post ID from the request path
	postID := c.Params("postId")

	// Retrieve the updated post data from the request body
	var updatedPost UpdatePostSchema
	if err := c.BodyParser(&updatedPost); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.FormatError(err.Error()))
	}

	// Validate the updated post data
	validate := validator.New()
	if err := validate.Struct(updatedPost); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.FormatError(err.Error()))
	}

	// Check if the post with the given ID exists
	existingPost, err := postsRepository.GetPost("id =?", postID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.FormatError(err.Error()))
	}
	// Update the post data
	updatePostFields(existingPost, updatedPost)
	fmt.Println(existingPost)
	// Save the updated post data
	newPost, err := postsRepository.UpdatePost(existingPost)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.FormatError(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(newPost)
}

// @Summary Delete post by ID
// @Tags Post
// @Accept json
// @Param postId path int false "Post ID"
// @Success 200 {object} string "Success"
// @Failure 400 {object} common.CommonErrorSchema "Bad Request"
// @Failure 404 {object} common.CommonErrorSchema "Not Found Request"
// @Router /v1/posts/{postId} [delete]
func deletePostById(c *fiber.Ctx) error {
	// Parse pagination parameters
	postID, err := parsePostIdParams(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.FormatError(err.Error()))
	}
	// Get post list based on query parameters
	post, err := postsRepository.GetPost("id =?", postID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(common.FormatError(POST_NOT_FOUND))
	}

	_, err = postsRepository.DeletePost(post)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.FormatError(err.Error()))
	}
	// Construct the response
	response := POST_SUCCESSFULLY_DELETED

	return c.Status(fiber.StatusOK).SendString(response)
}
