// user.go
package model

import "errors"

// User represents a user in the system
type User struct {
	ID       int    `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

var users = []User{
	{1, "user1", "password1"},
	{2, "user2", "password2"},
	{3, "user3", "password3"},
}

// GetUser retrieves a user from the database
func GetUser(login string) (User, error) {
	for _, user := range users {
		if user.Login == login {
			return user, nil
		}
	}
	return User{}, errors.New("User not found")
}
