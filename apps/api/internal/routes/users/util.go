package users

import (
	"dragon-threads/apps/api/internal/entities"
	"dragon-threads/apps/api/internal/repositories/usersRepository"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func isUserRegistered(email string, username string) bool {
	userByEmail, _ := usersRepository.GetUser("email = ?", email)
	userByUsername, _ := usersRepository.GetUser("username = ?", username)

	fmt.Println(userByUsername.ID, "ivan")
	// Check if either the email or username query found a user
	return userByEmail.ID > 0 || userByUsername.ID > 0
}

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
