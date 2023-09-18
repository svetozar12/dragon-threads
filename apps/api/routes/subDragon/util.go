package subDragon

import (
	"dragon-threads/apps/api/entities"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

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
	subDragonStr := c.Params("subDragon")
	subDragon, err := strconv.Atoi(subDragonStr)
	if err != nil {
		return 0, fmt.Errorf("SubDragon id is required")
	}
	return subDragon, nil
}

func updateSubDragonFields(existingSubDragon *entities.SubDragon, updatedSubDragon UpdateSubDragonSchema) {
	if updatedSubDragon.Name != existingSubDragon.Name {
		existingSubDragon.Name = updatedSubDragon.Name
	}

	if updatedSubDragon.Description != existingSubDragon.Description {
		existingSubDragon.Description = updatedSubDragon.Description
	}
}
