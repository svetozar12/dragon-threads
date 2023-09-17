package subDragon

import (
	"dragon-threads/apps/api/entities"
	"dragon-threads/apps/api/repositories/subDragonRepository"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func parsePaginationQuery(c *fiber.Ctx) (int, int) {
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

func parseGetByQuery(c *fiber.Ctx) (string, string) {
	getBy := c.Query("getBy")
	getByValue := c.Query("getByValue")
	if getBy == "" || getByValue == "" {
		getBy = "1"
		getByValue = "1"
	}

	return getBy, getByValue
}

func parseSubDragonIdParams(c *fiber.Ctx) (int, error) {
	userIdStr := c.Params("userId")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		return 0, fmt.Errorf("SubDragon id is required")
	}
	return userId, nil
}
func getSubDragonsByQuery(getBy string, getByValue string, page int, pageSize int) ([]entities.SubDragon, int64, error) {
	// Retrieve user list based on query parameters
	userList, total, err := subDragonRepository.GetSubDragonList(getBy+"=?", page, pageSize, []interface{}{getByValue})
	return userList, total, err
}

func updateSubDragonFields(existingSubDragon *entities.SubDragon, updatedSubDragon UpdateSubDragonSchema) {
	if updatedSubDragon.Name != existingSubDragon.Name {
		existingSubDragon.Name = updatedSubDragon.Name
	}

	if updatedSubDragon.Description != existingSubDragon.Description {
		existingSubDragon.Description = updatedSubDragon.Description
	}
}
