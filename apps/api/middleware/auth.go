package middleware

import (
	"dragon-threads/apps/api/pkg/common"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// validateToken validates the GitHub token
func validateToken(token string) bool {
	// Make a request to GitHub's token validation endpoint
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.github.com/user", nil)
	if err != nil {
		return false
	}

	req.Header.Set("Authorization", "Bearer "+token)
	resp, err := client.Do(req)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	// Check the response status
	if resp.StatusCode != http.StatusOK {
		return false
	}

	// You can optionally parse and use the response body if needed

	return true
}

func OAuth2Middleware(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(common.FormatError(fiber.ErrUnauthorized.Message))
	}

	// Validate the OAuth2 token (GitHub token in this case)
	if !validateToken(token) {
		return c.Status(fiber.StatusUnauthorized).JSON(common.FormatError(fiber.ErrUnauthorized.Message))
	}

	// Call the next middleware or handler
	return c.Next()
}
