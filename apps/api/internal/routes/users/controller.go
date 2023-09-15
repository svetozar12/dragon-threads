package users

import (
	"dragon-threads/apps/api/internal/entities"
	"dragon-threads/apps/api/internal/pkg/common"
	"dragon-threads/apps/api/internal/repositories/usersRepository"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// Example godoc
// @Summary      GEt USer List
// @Tags         User
// @Accept       json
// @Param page query int false "Page number (default: 1)"
// @Param pageSize query int false "Number of items per page (default: 10)"
// @Success      200  {object} users.UserListSchema "Success"
// @Failure 	 400  {object} common.CommonErrorSchema
// @Router       /v1/users [get]
func getUserList(c *fiber.Ctx) error {
	pageStr := c.Query("page")
	page, err := strconv.Atoi(pageStr)

	if err != nil {
		// If conversion to int fails, set a default value (1 in this case)
		page = 1
	}
	pageSizeStr := c.Query("pageSize")
	pageSize, err := strconv.Atoi(pageSizeStr)

	if err != nil {
		// If conversion to int fails, set a default value (1 in this case)
		pageSize = 10
	}
	userList, total, err := usersRepository.GetUserList("1=1", page, pageSize)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.FormatError(err.Error()))
	}
	hasNextPage := (page * pageSize) < int(total)
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
