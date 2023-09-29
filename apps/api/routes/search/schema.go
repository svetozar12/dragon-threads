package search

import (
	"dragon-threads/apps/api/entities"
	"dragon-threads/apps/api/pkg/common"
)

type SearchUsersSchema struct {
	Pagination common.CommonPaginationSchema `json:"pagination"`
	Data       []entities.User               `json:"data"`
}

type SearchSubDragonsSchema struct {
	Pagination common.CommonPaginationSchema `json:"pagination"`
	Data       []entities.SubDragon          `json:"data"`
}

type SearchPostsSchema struct {
	Pagination common.CommonPaginationSchema `json:"pagination"`
	Data       []entities.Post               `json:"data"`
}
