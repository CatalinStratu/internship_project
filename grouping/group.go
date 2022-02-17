package grouping

import (
	"awesomeProject2/user"
)

// Grouping all users by the first letter of Firstname field
func Grouping(users []user.User) map[string][]user.User {
	collections := make(map[string][]user.User)
	for _, u := range users {
		collections[u.FirstName[0:1]] = append(collections[u.FirstName[0:1]], u)
	}
	return collections
}
