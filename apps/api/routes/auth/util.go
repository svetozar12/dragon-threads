package auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type GithubEmail struct {
	Email      string `json:"email"`
	Primary    bool   `json:"primary"`
	Verified   bool   `json:"verified"`
	Visibility string `json:"visibility"`
}

func getPrimaryGithubEmail(emails []GithubEmail) (string, error) {
	for _, value := range emails {
		fmt.Println(value)
		fmt.Println(value.Email)
		if value.Primary {
			return value.Email, nil
		}
	}
	return "", fmt.Errorf("Your oauth email is invalid")
}

func getGithubUserEmail(client *http.Client) (string, error) {
	resp, err := client.Get("https://api.github.com/user/emails")
	if err != nil {
		return "", fmt.Errorf("Error fetching user details from GitHub")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return "", err
	}
	var data []GithubEmail
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return "", err
	}
	email, err := getPrimaryGithubEmail(data)
	return email, nil
}
