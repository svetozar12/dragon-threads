package subDragon

import (
	"dragon-threads/apps/api/constants"
	"dragon-threads/apps/api/entities"
	"dragon-threads/apps/api/pkg/common"
	"dragon-threads/apps/api/repositories/subDragonRepository"
	"dragon-threads/apps/api/repositories/usersRepository"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// @Summary Get subDragon by ID
// @Tags SubDragon
// @Accept json
// @Param subDragonId path int false "SubDragon ID"
// @Security ApiKeyAuth
// @Success 200 {object} entities.SubDragon "Success"
// @Failure 400 {object} common.CommonErrorSchema "Bad Request"
// @Failure 404 {object} common.CommonErrorSchema "Not Found Request"
// @Router /v1/subDragon/{subDragonId} [get]
func getSubDragonById(c *fiber.Ctx) error {
	// Parse pagination parameters
	subDragonID, err := parseSubDragonIdParams(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.FormatError(err.Error()))
	}
	// Get subDragon list based on query parameters
	subDragon, err := subDragonRepository.GetSubDragon("id =?", subDragonID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(common.FormatError(constants.SUB_DRAGON_NOT_FOUND))
	}
	// Construct the response
	response := subDragon

	return c.Status(fiber.StatusOK).JSON(response)
}

// @Summary      Get SubDragon List
// @Tags         SubDragon
// @Accept       json
// @Param page       query int false "Page number (default: 1)"
// @Param pageSize   query int false "Number of items per page (default: 10)"
// @Security ApiKeyAuth
// @Success      200  {object} subDragon.SubDragonListSchema "Success"
// @Failure 	 400  {object} common.CommonErrorSchema
// @Router       /v1/subDragon [get]
func getSubDragonList(c *fiber.Ctx) error {
	// Parse pagination parameters
	page, pageSize := common.ParsePaginationQuery(c)

	// Get subDragon list based on query parameters
	subDragonList, total, err := subDragonRepository.GetSubDragonList("1=1", page, pageSize, []interface{}{})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.FormatError(err.Error()))
	}

	// Check if there is a next page
	hasNextPage := (page * pageSize) < int(total)

	// Construct the response
	response := SubDragonListSchema{
		Data: subDragonList,
		Pagination: common.CommonPaginationSchema{
			Page:       page,
			PageSize:   pageSize,
			TotalCount: total,
			HasNext:    hasNextPage,
		},
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

// Example godoc
// @Summary      Create SubDragon
// @Tags         SubDragon
// @Accept       json
// @Param request body subDragon.CreateSubDragonSchema true "query params""
// @Security ApiKeyAuth
// @Success      201  {object} entities.SubDragon
// @Failure 	 400  {object} common.CommonErrorSchema
// @Router       /v1/subDragon [post]
func createSubDragon(c *fiber.Ctx) error {
	var subDragonPayload CreateSubDragonSchema
	if err := c.BodyParser(&subDragonPayload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.FormatError(err.Error()))
	}
	validate := validator.New()
	if err := validate.Struct(subDragonPayload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.FormatError(err.Error()))
	}

	_, err := subDragonRepository.GetSubDragon("name =?", subDragonPayload.Name)
	if err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.FormatError(constants.SUB_DRAGON_ALREADY_EXIST))
	}
	user, _ := usersRepository.GetUser("id =?", subDragonPayload.UserId)
	subDragon, err := subDragonRepository.CreateSubDragon(&entities.SubDragon{
		Name: subDragonPayload.Name,
	})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.FormatError(err.Error()))
	}
	user.SubDragonId = int32(subDragon.ID)
	usersRepository.UpdateUser(user)

	return c.Status(fiber.StatusOK).JSON(subDragon)
}

// @Summary      Update SubDragon
// @Tags         SubDragon
// @Accept       json
// @Param id path string true "SubDragon ID"
// @Param request body subDragon.UpdateSubDragonSchema true "Request body for updating subDragon"
// @Security ApiKeyAuth
// @Success      200  {object} entities.SubDragon
// @Failure      400  {object} common.CommonErrorSchema
// @Router       /v1/subDragon/{id} [put]
func updateSubDragon(c *fiber.Ctx) error {
	// Retrieve the subDragon ID from the request path
	subDragonID := c.Params("subDragonId")

	// Retrieve the updated subDragon data from the request body
	var updatedSubDragon UpdateSubDragonSchema
	if err := c.BodyParser(&updatedSubDragon); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.FormatError(err.Error()))
	}

	// Validate the updated subDragon data
	validate := validator.New()
	if err := validate.Struct(updatedSubDragon); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.FormatError(err.Error()))
	}

	// Check if the subDragon with the given ID exists
	existingSubDragon, err := subDragonRepository.GetSubDragon("id =?", subDragonID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.FormatError(err.Error()))
	}
	// Update the subDragon data
	updateSubDragonFields(existingSubDragon, updatedSubDragon)
	// Save the updated subDragon data
	newSubDragon, err := subDragonRepository.UpdateSubDragon(existingSubDragon)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.FormatError(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(newSubDragon)
}

// @Summary Delete subDragon by ID
// @Tags SubDragon
// @Accept json
// @Param subDragonId path int false "SubDragon ID"
// @Security ApiKeyAuth
// @Success 200 {object} string "Success"
// @Failure 400 {object} common.CommonErrorSchema "Bad Request"
// @Failure 404 {object} common.CommonErrorSchema "Not Found Request"
// @Router /v1/subDragon/{subDragonId} [delete]
func deleteSubDragonById(c *fiber.Ctx) error {
	// Parse pagination parameters
	subDragonID, err := parseSubDragonIdParams(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.FormatError(err.Error()))
	}
	// Get subDragon list based on query parameters
	subDragon, err := subDragonRepository.GetSubDragon("id =?", subDragonID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(common.FormatError(constants.SUB_DRAGON_NOT_FOUND))
	}

	_, err = subDragonRepository.DeleteSubDragon(subDragon)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.FormatError(err.Error()))
	}
	// Construct the response
	response := constants.SUB_DRAGON_SUCCESSFULLY_DELETED

	return c.Status(fiber.StatusOK).SendString(response)
}
