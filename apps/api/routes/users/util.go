package users

import (
	"dragon-threads/apps/api/entities"
	"dragon-threads/apps/api/pkg/common"
	"dragon-threads/apps/api/repositories/usersRepository"
	"fmt"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func isUserRegistered(email string, username string) *entities.User {
	userByEmail, _ := usersRepository.GetUser("email = ?", email)
	userByUsername, _ := usersRepository.GetUser("username = ?", username)

	fmt.Println(userByUsername.ID, "ivan")
	// Check if either the email or username query found a user
	if userByEmail != nil {
		return userByEmail
	} else if userByUsername != nil {
		return userByUsername
	}
	return nil
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

func parseUserIdParams(c *fiber.Ctx) (int, error) {
	userIdStr := c.Params("userId")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		return 0, fmt.Errorf("User id is required")
	}
	return userId, nil
}
func getUsersByQuery(getBy string, getByValue string, page int, pageSize int) ([]entities.User, int64, error) {
	// Retrieve user list based on query parameters
	userList, total, err := usersRepository.GetUserList(getBy+"=?", page, pageSize, []interface{}{getByValue})
	return userList, total, err
}

func updateUserFields(existingUser *entities.User, updatedUser UpdateUserSchema) {
	if updatedUser.Username != "" {
		existingUser.Username = updatedUser.Username
	}

	if updatedUser.Email != "" {
		existingUser.Email = updatedUser.Email
	}
}

func FindOrCreateUser(c *fiber.Ctx, user UserSchema) (*entities.User, error) {
	if err := c.BodyParser(&user); err != nil {
		return nil, c.Status(fiber.StatusBadRequest).JSON(common.FormatError(err.Error()))
	}
	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		return nil, c.Status(fiber.StatusBadRequest).JSON(common.FormatError(err.Error()))
	}
	if userInstance := isUserRegistered(user.Email, user.Username); userInstance != nil {
		return userInstance, nil
	}
	_, err := usersRepository.CreateUser(&entities.User{
		Username:    user.Username,
		Email:       user.Email,
		Bio:         user.Bio,
		Avatar:      user.Avatar,
		Active:      true,
		SubDragonId: user.SubDragonId,
	})
	if err != nil {
		return nil, c.Status(fiber.StatusBadRequest).JSON(common.FormatError(err.Error()))
	}
	return nil, c.Status(fiber.StatusOK).JSON(user)
}
