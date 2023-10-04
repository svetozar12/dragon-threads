package postComment

import (
	"dragon-threads/apps/api/constants"
	"dragon-threads/apps/api/entities"
	"dragon-threads/apps/api/pkg/common"
	"dragon-threads/apps/api/repositories/postCommentRepository"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// @Summary 			Get post by ID
// @Tags 				PostComment
// @Accept 				json
// @Param postId 		path int false "Post ID"
// @Param subDragonId	query int false "Search in SubDragonId"
// @Security 			ApiKeyAuth
// @Success 			200 {object} posts.PostSchema "Success"
// @Failure 			400 {object} common.CommonErrorSchema "Bad Request"
// @Failure 			404 {object} common.CommonErrorSchema "Not Found Request"
// @Router 			 	/v1/posts/comment/{postId} [get]
func GetPostCommentById(c *fiber.Ctx) error {
	// Parse pagination parameters
	postCommentID := common.ParseIntParam(c, constants.POST_COMMENT_ID, 0)
	postID := common.ParseIntParam(c, constants.POST_ID, 0)
	userID := common.ParseIntParam(c, constants.USER_ID, 0)

	// Get post list based on query parameters
	postComment, err := postCommentRepository.GetPostComment("id =? AND post_id =? AND user_id =?", postCommentID, postID, userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(common.FormatError(constants.POST_COMMENT_NOT_FOUND))
	}
	// Construct the response
	response := postComment

	return c.Status(fiber.StatusOK).JSON(response)
}

// @Summary      		Get Post List
// @Tags         		PostComment
// @Accept       		json
// @Param page			query int false "Page number (default: 1)"
// @Param pageSize		query int false "Number of items per page (default: 10)"
// @Param postId		query int false "Search by post id"
// @Security 			ApiKeyAuth
// @Success      		200  {object} posts.PostListSchema "Success"
// @Failure 	 		400  {object} common.CommonErrorSchema
// @Router       		/v1/posts/comment [get]
func getPostCommentList(c *fiber.Ctx) error {
	// Parse pagination parameters
	page, pageSize := common.ParsePaginationQuery(c)
	postID := common.ParseIntParam(c, constants.POST_ID, 0)

	prePost := c.Locals(common.POST_BY_ID(postID)).(entities.PostComment)

	// Get post list based on query parameters
	postList, total, err := getPostCommentListByQuery(page, pageSize, int(prePost.ID))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.FormatError(err.Error()))
	}

	// Check if there is a next page
	hasNextPage := (page * pageSize) < int(total)

	// Construct the response
	response := PostCommentListSchema{
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

// @Summary      		Create PostComment
// @Tags         		PostComment
// @Accept       		json
// @Param request body  posts.PostSchema true "query params""
// @Security 			ApiKeyAuth
// @Success      		201  {object} entities.PostComment
// @Failure 	 		400  {object} common.CommonErrorSchema
// @Router       		/v1/posts/comment [post]
func createPostComment(c *fiber.Ctx) error {
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

// @Summary      		Update PostComment
// @Tags         		PostComment
// @Accept       		json
// @Param id 			path string true "PostComment ID"
// @Param subDragonId	query int false "Search in SubDragonId"
// @Param request body 	posts.UpdatePostSchema true "Request body for updating post"
// @Security 			ApiKeyAuth
// @Success      		200  {object} entities.PostComment
// @Failure      		400  {object} common.CommonErrorSchema
// @Router       		/v1/posts/comment/{id} [put]
func updatePostComment(c *fiber.Ctx) error {
	// Retrieve the post ID from the request path
	postID := common.ParseIntParam(c, constants.POST_ID, 0)
	preSubDragon := c.Locals(common.SUB_DRAGON_BY_ID).(entities.SubDragon)
	preUser := c.Locals(common.USER_BY_ID).(entities.User)
	// Retrieve the updated post data from the request body
	var updatedPost UpdatePostCommentSchema
	if err := c.BodyParser(&updatedPost); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.FormatError(err.Error()))
	}

	// Validate the updated post data
	validate := validator.New()
	if err := validate.Struct(updatedPost); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.FormatError(err.Error()))
	}

	// Check if the post with the given ID exists
	existingPost, err := postCommentRepository.GetPostComment("id =? AND post_id=?", postID, preUser.ID, preSubDragon.ID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.FormatError(err.Error()))
	}
	// Update the post data
	updatePostCommentFields(existingPost, updatedPost)
	// Save the updated post data
	newPost, err := postCommentRepository.UpdatePostComment(existingPost)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.FormatError(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(newPost)
}

// @Summary 			Delete post by ID
// @Tags 				PostComment
// @Accept 				json
// @Param postId 		path int false "Post ID"
// @Param subDragonId	query int false "Search in SubDragonId"
// @Param userId		query int false "Search in UserId"
// @Security 			ApiKeyAuth
// @Success 			200 {object} string "Success"
// @Failure 			400 {object} common.CommonErrorSchema "Bad Request"
// @Failure 			404 {object} common.CommonErrorSchema "Not Found Request"
// @Router 				/v1/posts/comment/{postId} [delete]
func deletePostCommentById(c *fiber.Ctx) error {
	// Parse pagination parameters
	postCommentID := common.ParseIntParam(c, constants.POST_COMMENT_ID, 0)
	postID := common.ParseIntParam(c, constants.POST_ID, 0)
	userID := common.ParseIntParam(c, constants.USER_ID, 0)
	// Get post list based on query parameters
	postComment, err := postCommentRepository.GetPostComment("id =? AND post_id =? AND user_id =?", postCommentID, postID, userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(common.FormatError(constants.POST_COMMENT_NOT_FOUND))
	}

	_, err = postCommentRepository.DeletePostComment(postComment)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.FormatError(err.Error()))
	}
	// Construct the response
	response := constants.POST_SUCCESSFULLY_DELETED

	return c.Status(fiber.StatusOK).SendString(response)
}
