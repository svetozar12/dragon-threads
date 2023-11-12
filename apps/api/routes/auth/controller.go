package auth

import (
	"dragon-threads/apps/api/definitions"
	"dragon-threads/apps/api/pkg/env"
	"dragon-threads/apps/api/pkg/oauth"
	"dragon-threads/apps/api/routes/users"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/oauth2"
)

// @Summary Login with github
// @Tags Auth
// @Router /v1/auth/github [get]
func login(c *fiber.Ctx) error {
	url := oauth.GithubOauthConfig.AuthCodeURL("state", oauth2.AccessTypeOnline)
	return c.Redirect(url)
}

// @Summary Callback for corresponding login method
// @Tags Auth
// @Success 200 {object} auth.AuthSchema "Success"
// @Router /v1/auth/github/callback [get]
func githubCallback(c *fiber.Ctx) error {
	code := c.Query("code")
	token, err := oauth.GithubOauthConfig.Exchange(c.Context(), code)
	if err != nil {
		return c.Redirect(env.Envs.FRONTEND_URL)
	}
	client := oauth.GithubOauthConfig.Client(c.Context(), token)
	resp, err := client.Get("https://api.github.com/user")
	if err != nil {
		return c.Redirect(env.Envs.FRONTEND_URL)
	}
	var user definitions.GithubUser
	// Read the response body into a byte slice
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return c.Redirect(env.Envs.FRONTEND_URL)
	}
	// Unmarshal the JSON data into the GithubUser struct
	err = json.Unmarshal(body, &user)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return c.Redirect(env.Envs.FRONTEND_URL)
	}
	email, err := getGithubUserEmail(client)
	// Assuming you have a function to find or create the user
	// Replace with your actual logic for finding or creating the user
	_, err = users.FindOrCreateUser(c, users.UserSchema{Username: user.Login, Email: email, Avatar: user.AvatarURL, Bio: user.Bio})
	if err != nil {
		return c.Redirect(env.Envs.FRONTEND_URL)
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

// Example godoc
// @Summary      Verify token
// @Tags         Auth
// @Accept       json
// @Security 	 ApiKeyAuth
// @Success      201  {object} nil
// @Failure 	 400  {object} common.CommonErrorSchema
// @Router       /v1/auth/verify [get]
func verifyToken(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(nil)
}
