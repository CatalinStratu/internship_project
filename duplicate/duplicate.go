package duplicate

import (
	api "awesomeProject2/service"
)

// Duplicate interface
type Duplicate interface {
	Remove(resSlice []api.User) []api.User
}

// Remove duplicates
func Remove(users []api.User) []api.User {
	allKeys := make(map[api.User]bool)
	var list []api.User
	for _, item := range users {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}
