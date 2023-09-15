package users

import (
	"dragon-threads/apps/api/internal/entities"
	"dragon-threads/apps/api/internal/pkg/common"
)

type UserSchema struct {
	Username    string `json:"username"`
	Email       string `json:"email"`
	Avatar      string `json:"avatar"`
	Bio         string `json:"bio"`
	SubDragonId int32  `json:"subDragonId"`
}

type UserListSchema struct {
	Pagination common.CommonPaginationSchema `json:"pagination"`
	Data       []entities.User               `json:"data"`
}
