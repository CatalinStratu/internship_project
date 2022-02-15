package grouping

import api "awesomeProject2/service"

func Grouping(users []api.User) map[string][]api.User {
	collections := make(map[string][]api.User)
	for _, user := range users {
		collections[user.FirstName[0:1]] = append(collections[user.FirstName[0:1]], user)
	}
	return collections
}
