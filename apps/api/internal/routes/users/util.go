package users

import (
	"dragon-threads/apps/api/internal/repositories/usersRepository"
	"fmt"
)

func isUserRegistered(email string, username string) bool {
	userByEmail, _ := usersRepository.GetUser("email = ?", email)
	userByUsername, _ := usersRepository.GetUser("username = ?", username)

	fmt.Println(userByUsername.ID, "ivan")
	// Check if either the email or username query found a user
	return userByEmail.ID > 0 || userByUsername.ID > 0
}
