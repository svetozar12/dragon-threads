package search

import (
	"dragon-threads/apps/api/constants"
	"dragon-threads/apps/api/pkg/common"
	"dragon-threads/apps/api/repositories/postsRepository"
	"dragon-threads/apps/api/repositories/subDragonRepository"
	"dragon-threads/apps/api/repositories/usersRepository"

	"github.com/gofiber/fiber/v2"
)

// @Summary Search users by username
// @Tags     Search
// @Accept   json
// @Param    searchText query string false "Search by username"
// @Param page       query int false "Page number (default: 1)"
// @Param pageSize   query int false "Number of items per page (default: 10)"
// @Security ApiKeyAuth
// @Success  200 {object} search.SearchUsersSchema "Success"
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

// @Summary Search subDragons by name
// @Tags     Search
// @Accept   json
// @Param    searchText query string false "Search by subDragon name"
// @Param page       query int false "Page number (default: 1)"
// @Param pageSize   query int false "Number of items per page (default: 10)"
// @Security ApiKeyAuth
// @Success  200 {object} search.SearchSubDragonsSchema "Success"
// @Router /v1/search/subDragon [get]
func SearchSubDragons(c *fiber.Ctx) error {
	// Parse pagination parameters
	page, pageSize := common.ParsePaginationQuery(c)
	searchText := c.Query("searchText")
	// Get user list based on query parameters
	subDragons, total, _ := subDragonRepository.SearchSubDragonByText(searchText, page, pageSize)
	// Check if there is a next page
	hasNextPage := (page * pageSize) < int(total)

	return c.Status(fiber.StatusOK).JSON(&SearchSubDragonsSchema{
		Data: subDragons,
		Pagination: common.CommonPaginationSchema{
			Page:       page,
			PageSize:   pageSize,
			TotalCount: total,
			HasNext:    hasNextPage,
		}})
}

// @Summary Search posts by post name
// @Tags     Search
// @Accept   json
// @Param    searchText query string false "Search by subDragon name"
// @Param page       query int false "Page number (default: 1)"
// @Param pageSize   query int false "Number of items per page (default: 10)"
// @Param subDragonId	query int false "Search in SubDragonId"
// @Security ApiKeyAuth
// @Success  200 {object} search.SearchSubDragonsSchema "Success"
// @Router /v1/search/posts [get]
func SearchPosts(c *fiber.Ctx) error {
	// Parse pagination parameters
	page, pageSize := common.ParsePaginationQuery(c)
	searchText := c.Query("searchText")
	subDragonId := common.ParseIntParam(c, constants.SUB_DRAGON_ID, 0)
	// Get user list based on query parameters
	posts, total, _ := postsRepository.SearchPostByText(searchText, page, pageSize, subDragonId)
	// Check if there is a next page
	hasNextPage := (page * pageSize) < int(total)

	return c.Status(fiber.StatusOK).JSON(&SearchPostsSchema{
		Data: posts,
		Pagination: common.CommonPaginationSchema{
			Page:       page,
			PageSize:   pageSize,
			TotalCount: total,
			HasNext:    hasNextPage,
		}})
}
