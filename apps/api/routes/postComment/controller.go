package postComment

import (
	"dragon-threads/apps/api/constants"
	"dragon-threads/apps/api/entities"
	"dragon-threads/apps/api/pkg/common"
	"dragon-threads/apps/api/repositories/postCommentRepository"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// @Summary Get post by ID
// @Tags PostComment
// @Accept json
// @Param postId path int false "Post ID"
// @Param subDragonId	query int false "Search in SubDragonId"
// @Security ApiKeyAuth
// @Success 200 {object} posts.PostSchema "Success"
// @Failure 400 {object} common.CommonErrorSchema "Bad Request"
// @Failure 404 {object} common.CommonErrorSchema "Not Found Request"
// @Router /v1/posts/comment/{postId} [get]
func GetPostById(c *fiber.Ctx) error {
	// Parse pagination parameters
	postID := common.ParseIntParam(c, constants.POST_ID, 0)
	preSubDragon := c.Locals(common.SUB_DRAGON_BY_ID).(entities.SubDragon)
	preUser := c.Locals(common.USER_BY_ID).(entities.User)

	// Get post list based on query parameters
	post, err := postCommentRepository.GetPost("id =? AND sub_dragon_id =? AND user_id =?", postID, preSubDragon.ID, preUser.ID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(common.FormatError(constants.POST_NOT_FOUND))
	}
	// Construct the response
	response := post

	return c.Status(fiber.StatusOK).JSON(response)
}

// @Summary      Get Post List
// @Tags         Post
// @Accept       json
// @Param page			query int false "Page number (default: 1)"
// @Param pageSize		query int false "Number of items per page (default: 10)"
// @Param subDragonId	query int false "Search in SubDragonId"
// @Security ApiKeyAuth
// @Success      200  {object} posts.PostListSchema "Success"
// @Failure 	 400  {object} common.CommonErrorSchema
// @Router       /v1/posts/comment [get]
func getPostList(c *fiber.Ctx) error {
	// Parse pagination parameters
	page, pageSize := common.ParsePaginationQuery(c)
	preSubDragon := c.Locals(common.SUB_DRAGON_BY_ID).(entities.SubDragon)
	preUser := c.Locals(common.USER_BY_ID).(entities.User)

	// Get post list based on query parameters
	postList, total, err := getPostsByQuery(page, pageSize, int(preSubDragon.ID), int(preUser.ID))
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
// @Security ApiKeyAuth
// @Success      201  {object} entities.Post
// @Failure 	 400  {object} common.CommonErrorSchema
// @Router       /v1/posts/comment [post]
func createPost(c *fiber.Ctx) error {
	var postComment PostCommentSchema
	if err := c.BodyParser(&postComment); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.FormatError(err.Error()))
	}
	validate := validator.New()
	if err := validate.Struct(postComment); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.FormatError(err.Error()))
	}

	_, err := postCommentRepository.CreatePostComment(&entities.PostComment{
		UserID:  postComment.UserID,
		Comment: postComment.Comment,
		PostId:  postComment.PostID,
	})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.FormatError(err.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(postComment)
}

// @Summary      Update Post
// @Tags         Post
// @Accept       json
// @Param id path string true "Post ID"
// @Param subDragonId	query int false "Search in SubDragonId"
// @Param request body posts.UpdatePostSchema true "Request body for updating post"
// @Security ApiKeyAuth
// @Success      200  {object} entities.Post
// @Failure      400  {object} common.CommonErrorSchema
// @Router       /v1/posts/comment/{id} [put]
func updatePost(c *fiber.Ctx) error {
	// Retrieve the post ID from the request path
	postID := common.ParseIntParam(c, constants.POST_ID, 0)
	preSubDragon := c.Locals(common.SUB_DRAGON_BY_ID).(entities.SubDragon)
	preUser := c.Locals(common.USER_BY_ID).(entities.User)
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
	existingPost, err := postCommentRepository.GetPost("id =? AND user_id=? AND sub_dragon_id=?", postID, preUser.ID, preSubDragon.ID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.FormatError(err.Error()))
	}
	// Update the post data
	updatePostFields(existingPost, updatedPost)
	// Save the updated post data
	newPost, err := postCommentRepository.UpdatePost(existingPost)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.FormatError(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(newPost)
}

// @Summary Delete post by ID
// @Tags PostComment
// @Accept json
// @Param postId path int false "Post ID"
// @Param subDragonId	query int false "Search in SubDragonId"
// @Param userId		query int false "Search in UserId"
// @Security ApiKeyAuth
// @Success 200 {object} string "Success"
// @Failure 400 {object} common.CommonErrorSchema "Bad Request"
// @Failure 404 {object} common.CommonErrorSchema "Not Found Request"
// @Router /v1/posts/comment/{postId} [delete]
func deletePostById(c *fiber.Ctx) error {
	// Parse pagination parameters
	postID := common.ParseIntParam(c, constants.POST_ID, 0)
	preSubDragon := c.Locals(common.SUB_DRAGON_BY_ID).(entities.SubDragon)
	preUser := c.Locals(common.USER_BY_ID).(entities.User)

	// Get post list based on query parameters
	post, err := postCommentRepository.GetPost("id =? AND user_id =? AND sub_dragon_id =?", postID, preUser.ID, preSubDragon.ID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(common.FormatError(constants.POST_NOT_FOUND))
	}

	_, err = postCommentRepository.DeletePost(post)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.FormatError(err.Error()))
	}
	// Construct the response
	response := constants.POST_SUCCESSFULLY_DELETED

	return c.Status(fiber.StatusOK).SendString(response)
}
