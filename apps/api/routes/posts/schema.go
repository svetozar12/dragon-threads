package posts

import (
	"dragon-threads/apps/api/entities"
	"dragon-threads/apps/api/pkg/common"
)

type PostSchema struct {
	Title       string `json:"title"`
	Content     string `json:"content"`
	UserID      uint   `json:"user_id"` // Foreign key to User
	SubDragonId int32  `json:"subDragonId"`
}

type UpdatePostSchema struct {
	Title       string `json:"title"`
	Content     string `json:"content"`
	UserID      uint   `json:"user_id"` // Foreign key to User
	SubDragonId int32  `json:"subDragonId"`
}

type PostListSchema struct {
	Pagination common.CommonPaginationSchema `json:"pagination"`
	Data       []entities.Post               `json:"data"`
}
