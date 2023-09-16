package users

import (
	"dragon-threads/apps/api/internal/entities"
	"dragon-threads/apps/api/internal/pkg/common"
	"dragon-threads/apps/api/internal/repositories/usersRepository"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type GetByEnum string

const (
	SubDragonId GetByEnum = "sub_dragon_id"
)

// @Summary Get user by ID
// @Tags User
// @Accept json
// @Param userId path int false "User ID"
// @Success 200 {object} users.UserSchema "Success"
// @Failure 400 {object} common.CommonErrorSchema "Bad Request"
// @Failure 404 {object} common.CommonErrorSchema "Not Found Request"
// @Router /v1/users/{userId} [get]
func getUserById(c *fiber.Ctx) error {
	// Parse pagination parameters
	userID, err := parseUserIdParams(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.FormatError(err.Error()))
	}
	// Get user list based on query parameters
	user, err := usersRepository.GetUser("id =?", userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(common.FormatError(USER_NOT_FOUND))
	}
	// Construct the response
	response := user

	return c.Status(fiber.StatusOK).JSON(response)
}

// @Summary      Get User List
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
	page, pageSize := parsePaginationQuery(c)

	// Get query and value parameters
	getBy, getByValue := parseGetByQuery(c)

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

// @Summary      Update User
// @Tags         User
// @Accept       json
// @Param id path string true "User ID"
// @Param request body users.UpdateUserSchema true "Request body for updating user"
// @Success      200  {object} entities.User
// @Failure      400  {object} common.CommonErrorSchema
// @Router       /v1/users/{id} [put]
func updateUser(c *fiber.Ctx) error {
	// Retrieve the user ID from the request path
	userID := c.Params("userId")

	// Retrieve the updated user data from the request body
	var updatedUser UpdateUserSchema
	if err := c.BodyParser(&updatedUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.FormatError(err.Error()))
	}

	// Validate the updated user data
	validate := validator.New()
	if err := validate.Struct(updatedUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.FormatError(err.Error()))
	}

	// Check if the user with the given ID exists
	existingUser, err := usersRepository.GetUser("id =?", userID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.FormatError(err.Error()))
	}
	// Update the user data
	updateUserFields(existingUser, updatedUser)
	fmt.Println(existingUser)
	// Save the updated user data
	newUser, err := usersRepository.UpdateUser(existingUser)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.FormatError(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(newUser)
}

// @Summary Delete user by ID
// @Tags User
// @Accept json
// @Param userId path int false "User ID"
// @Success 200 {object} string "Success"
// @Failure 400 {object} common.CommonErrorSchema "Bad Request"
// @Failure 404 {object} common.CommonErrorSchema "Not Found Request"
// @Router /v1/users/{userId} [delete]
func deleteUserById(c *fiber.Ctx) error {
	// Parse pagination parameters
	userID, err := parseUserIdParams(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.FormatError(err.Error()))
	}
	// Get user list based on query parameters
	user, err := usersRepository.GetUser("id =?", userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(common.FormatError(USER_NOT_FOUND))
	}

	_, err = usersRepository.DeleteUser(user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.FormatError(err.Error()))
	}
	// Construct the response
	response := USER_SUCCESSFULLY_DELETED

	return c.Status(fiber.StatusOK).SendString(response)
}
