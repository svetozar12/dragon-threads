package oauth

import (
	"dragon-threads/apps/api/pkg/env"
	"fmt"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

var (
	GithubOauthConfig *oauth2.Config
)

func InitGithubOauth() {
	fmt.Println(env.Envs)
	GithubOauthConfig = &oauth2.Config{
		ClientID:     env.Envs.GITHUB_CLIENT_ID,
		ClientSecret: env.Envs.GITHUB_SECRET,
		RedirectURL:  env.Envs.GITHUB_REDIRECT_URL,
		Endpoint:     github.Endpoint,
	}
}
