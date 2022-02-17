package grouping

import (
	"awesomeProject2/user"
)

func Grouping(users []user.User) map[string][]user.User {
	collections := make(map[string][]user.User)
	for _, user := range users {
		collections[user.FirstName[0:1]] = append(collections[user.FirstName[0:1]], user)
	}
	return collections
}
