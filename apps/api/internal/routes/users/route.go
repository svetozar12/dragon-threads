package users

import (
	"dragon-threads/apps/api/internal/entities"
	"dragon-threads/apps/api/internal/repositories/usersRepository"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func UsersRoute(app fiber.Router) {
	Users := app.Group("/users")
	Users.Post("/", createUser)
}

// Example godoc
// @Summary      Create User
// @Tags         User
// @Accept       json
// @Param request body users.UserSchema true "query params""
// @Success      201  {object} entities.User
// @Router       /v1/users [post]
func createUser(c *fiber.Ctx) error {
	var user entities.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(formatError(err.Error()))
	}
	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(formatError(err.Error()))
	}
	if isRegistered := isUserRegistered(user.Email, user.Username); isRegistered {
		return c.Status(fiber.StatusConflict).JSON(formatError(USER_ALREADY_EXIST))

	}
	user.Active = true
	_, err := usersRepository.CreateUser(&user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(formatError(err.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(user)
}
