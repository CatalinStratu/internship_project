package grouping

import (
	"service/user"
)

// Execute grouping all users by the first letter of Firstname field
func Execute(users []user.User) map[string][]user.User {
	collections := make(map[string][]user.User)
	for _, u := range users {
		if len(u.FirstName) != 0 {
			collections[u.FirstName[0:1]] = append(collections[u.FirstName[0:1]], u)
		}
	}
	return collections
}
