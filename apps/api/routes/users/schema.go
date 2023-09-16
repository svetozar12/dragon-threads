package users

import (
	"dragon-threads/apps/api/entities"
	"dragon-threads/apps/api/pkg/common"
)

type UserSchema struct {
	Username    string `json:"username" validate:"required,min=3,max=30"`
	Email       string `json:"email" validate:"required,email"`
	Avatar      string `json:"avatar"`
	Bio         string `json:"bio"`
	SubDragonId int32  `json:"subDragonId"`
}

type UpdateUserSchema struct {
	Username    string `json:"username" validate:"omitempty,min=3,max=30"`
	Email       string `json:"email" validate:"omitempty,email"`
	Avatar      string `json:"avatar" validate:"omitempty"`
	Bio         string `json:"bio" validate:"omitempty"`
	SubDragonId int32  `json:"subDragonId" validate:"omitempty"`
}

type UserListSchema struct {
	Pagination common.CommonPaginationSchema `json:"pagination"`
	Data       []entities.User               `json:"data"`
}
