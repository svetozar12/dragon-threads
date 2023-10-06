package postComment

import (
	"dragon-threads/apps/api/entities"
	"dragon-threads/apps/api/repositories/postCommentRepository"
)

func getPostCommentListByQuery(page int, pageSize int, postID int) ([]entities.PostComment, int64, error) {
	// Retrieve user list based on query parameters
	postList, total, err := postCommentRepository.GetPostCommentList("post_id=?", page, pageSize, []interface{}{postID})
	return postList, total, err
}

func updatePostCommentFields(existingPost *entities.PostComment, updatedPost UpdatePostCommentSchema) {
	if updatedPost.Comment != existingPost.Comment {
		existingPost.Comment = updatedPost.Comment
	}
}
