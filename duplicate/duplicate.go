package duplicate

import (
	"service/user"
)

// Remove duplicates
func Remove(users []user.User) []user.User {
	allKeys := make(map[user.User]struct{})
	var list []user.User
	for _, item := range users {
		if _, ok := allKeys[item]; !ok {
			allKeys[item] = struct{}{}
			list = append(list, item)
		}
	}
	return list
}
