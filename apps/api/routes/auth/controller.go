package auth

import (
	"dragon-threads/apps/api/pkg/common"
	"dragon-threads/apps/api/pkg/env"
	"dragon-threads/apps/api/pkg/oauth"

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
	// Set the access token as a cookie
	c.Cookie(&fiber.Cookie{
		Name:   "accessToken",
		Value:  token.AccessToken,
		MaxAge: 3600, // Set the cookie to expire in 1 hour (adjust as needed)
	})

	// Redirect the user to a URL
	return c.Redirect(env.Envs.FRONTEND_URL)
}
