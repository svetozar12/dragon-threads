package common

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func ParsePaginationQuery(c *fiber.Ctx) (int, int) {
	pageStr := c.Query("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}

	pageSizeStr := c.Query("pageSize")
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		pageSize = 10
	}

	return page, pageSize
}
