package common

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func ParsePaginationQuery(c *fiber.Ctx) (int, int) {
	page := ParseIntQuery(c, "page", 1)
	pageSize := ParseIntQuery(c, "pageSize", 10)

	return page, pageSize
}

func ParseIntQuery(c *fiber.Ctx, queryName string, defaultValue int) int {
	queryStr := c.Query(queryName)
	query, err := strconv.Atoi(queryStr)
	if err != nil {
		query = defaultValue
	}
	return query
}

func ParseIntParam(c *fiber.Ctx, queryName string, defaultValue int) int {
	queryStr := c.Params(queryName)
	query, err := strconv.Atoi(queryStr)
	if err != nil {
		query = defaultValue
	}
	return query
}
