package duplicate

import (
	"awesomeProject2/user"
)

// Remove duplicates
func Remove(users []user.User) []user.User {
	allKeys := make(map[user.User]bool)
	var list []user.User
	for _, item := range users {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}
