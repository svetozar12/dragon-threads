package users

import (
	"dragon-threads/apps/api/internal/entities"
	"dragon-threads/apps/api/internal/pkg/common"
	"dragon-threads/apps/api/internal/repositories/usersRepository"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type GetByEnum string

const (
	SubDragonId GetByEnum = "sub_dragon_id"
)

// Example godoc
// @Summary      GEt USer List
// @Tags         User
// @Accept       json
// @Param page       query int false "Page number (default: 1)"
// @Param pageSize   query int false "Number of items per page (default: 10)"
// @Param getBy      query GetByEnum false "Get users by field (optional)"
// @Param getByValue query int false "Get users by field value (optional)".
// @Success      200  {object} users.UserListSchema "Success"
// @Failure 	 400  {object} common.CommonErrorSchema
// @Router       /v1/users [get]
func getUserList(c *fiber.Ctx) error {
	// Parse pagination parameters
	page, pageSize := parsePaginationParams(c)

	// Get query and value parameters
	getBy, getByValue := parseQueryParams(c)

	// Get user list based on query parameters
	userList, total, err := getUsersByQuery(getBy, getByValue, page, pageSize)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.FormatError(err.Error()))
	}

	// Check if there is a next page
	hasNextPage := (page * pageSize) < int(total)

	// Construct the response
	response := UserListSchema{
		Data: userList,
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
// @Summary      Create User
// @Tags         User
// @Accept       json
// @Param request body users.UserSchema true "query params""
// @Success      201  {object} entities.User
// @Failure 	 400  {object} common.CommonErrorSchema
// @Router       /v1/users [post]
func createUser(c *fiber.Ctx) error {
	var user entities.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.FormatError(err.Error()))
	}
	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.FormatError(err.Error()))
	}
	if isRegistered := isUserRegistered(user.Email, user.Username); isRegistered {
		return c.Status(fiber.StatusConflict).JSON(common.FormatError(USER_ALREADY_EXIST))

	}
	user.Active = true
	_, err := usersRepository.CreateUser(&user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.FormatError(err.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(user)
}
