package subDragon

import (
	"dragon-threads/apps/api/entities"
	"dragon-threads/apps/api/pkg/common"
)

type CreateSubDragonSchema struct {
	Name   string `json:"title"`
	UserId int32  `json:"userId"`
}

type UpdateSubDragonSchema struct {
	Name            string   `json:"title" validate:"omitempty"`
	UserId          int32    `json:"userId" validate:"required"`
	Description     string   `json:"description" validate:"omitempty,min=1,max=500"`
	CommunityTopics []string `json:"communityTopics" validate:"omitempty"`
}

type SubDragonListSchema struct {
	Pagination common.CommonPaginationSchema `json:"pagination"`
	Data       []entities.SubDragon          `json:"data"`
}
