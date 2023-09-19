package auth

import (
	"dragon-threads/apps/api/definitions"
	"dragon-threads/apps/api/entities"
	"dragon-threads/apps/api/pkg/common"
	"dragon-threads/apps/api/pkg/env"
	"dragon-threads/apps/api/pkg/oauth"
	"dragon-threads/apps/api/routes/users"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/oauth2"
)

// @Summary Login with github
// @Tags Auth
// @Router /v1/auth/login [get]
func login(c *fiber.Ctx) error {
	url := oauth.GithubOauthConfig.AuthCodeURL("state", oauth2.AccessTypeOnline)
	return c.Redirect(url)
}

// @Summary Callback for corresponding login method
// @Tags Auth
// @Success 200 {object} auth.AuthSchema "Success"
// @Router /v1/auth/callback [get]
func githubCallback(c *fiber.Ctx) error {
	code := c.Query("code")
	token, err := oauth.GithubOauthConfig.Exchange(c.Context(), code)
	if err != nil {
		return c.JSON(common.FormatError("Error exchanging code for token"))
	}
	client := oauth.GithubOauthConfig.Client(c.Context(), token)
	resp, err := client.Get("https://api.github.com/user")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error fetching user details from GitHub",
		})
	}
	defer resp.Body.Close()

	var userDetail map[string]definitions.GithubUser
	if err := json.NewDecoder(resp.Body).Decode(&userDetail); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error decoding user details response",
		})
	}
	preSubDragon := c.Locals(common.SUB_DRAGON_BY_ID).(entities.SubDragon)
	// Assuming you have a function to find or create the user
	// Replace with your actual logic for finding or creating the user
	for _, user := range userDetail {
		users.FindOrCreateUser(c, users.UserSchema{Username: user.Login, Email: user.Email, Avatar: user.AvatarURL, Bio: user.Bio, SubDragonId: int32(preSubDragon.ID)})
	}
	// Set the access token as a cookie
	c.Cookie(&fiber.Cookie{
		Name:   "accessToken",
		Value:  token.AccessToken,
		MaxAge: 3600, // Set the cookie to expire in 1 hour (adjust as needed)
	})

	// Redirect the user to a URL
	return c.Redirect(env.Envs.FRONTEND_URL)
}
