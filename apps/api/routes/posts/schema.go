package posts

import (
	"dragon-threads/apps/api/entities"
	"dragon-threads/apps/api/pkg/common"
)

type PostSchema struct {
	Title       string `json:"title" validate:"required,min=3,max=20"`
	Content     string `json:"content" validate:"required,min=1,max=300"`
	UserID      int32  `json:"user_id" validate:"required"` // Foreign key to User
	SubDragonId int32  `json:"subDragonId" validate:"required"`
}

type UpdatePostSchema struct {
	Title   string `json:"title" validate:"omitempty"`
	Content string `json:"content" validate:"omitempty"`
}

type PostListSchema struct {
	Pagination common.CommonPaginationSchema `json:"pagination"`
	Data       []entities.Post               `json:"data"`
}
