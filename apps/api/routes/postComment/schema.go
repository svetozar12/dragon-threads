package postComment

import (
	"dragon-threads/apps/api/entities"
	"dragon-threads/apps/api/pkg/common"
)

type PostCommentSchema struct {
	Comment string `json:"comment" validate:"required,min=3,max=20"`
	UserID  int32  `json:"user_id" validate:"required"` // Foreign key to User
	PostID  int32  `json:"postId" validate:"required"`  // Foreign key to Post
}

type UpdatePostCommentSchema struct {
	Comment string `json:"comment" validate:"omitempty"`
}

type PostCommentListSchema struct {
	Pagination common.CommonPaginationSchema `json:"pagination"`
	Data       []entities.PostComment        `json:"data"`
}
