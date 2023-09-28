package search

import (
	"dragon-threads/apps/api/pkg/common"
	"dragon-threads/apps/api/repositories/usersRepository"

	"github.com/gofiber/fiber/v2"
)

type GetByEnum string

const (
	SubDragonId GetByEnum = "sub_dragon_id"
)

// @Summary Search users by
// @Tags     Search
// @Accept   json
// @Param    searchText query string false "Search by username"
// @Param page       query int false "Page number (default: 1)"
// @Param pageSize   query int false "Number of items per page (default: 10)"
// @Security ApiKeyAuth
// @Success  200 {object} users.UserSchema "Success"
// @Failure  400 {object} common.CommonErrorSchema "Bad Request"
// @Router /v1/search/users [get]
func SearchUsers(c *fiber.Ctx) error {
	// Parse pagination parameters
	page, pageSize := common.ParsePaginationQuery(c)
	searchText := c.Query("searchText")
	// Get user list based on query parameters
	users, total, _ := usersRepository.SearchUsersByText(searchText, page, pageSize)
	// Check if there is a next page
	hasNextPage := (page * pageSize) < int(total)

	return c.Status(fiber.StatusOK).JSON(&SearchUsersSchema{
		Data: users,
		Pagination: common.CommonPaginationSchema{
			Page:       page,
			PageSize:   pageSize,
			TotalCount: total,
			HasNext:    hasNextPage,
		}})
}
