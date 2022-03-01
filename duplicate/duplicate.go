package duplicate

import (
	"service/user"
)

// Remove duplicates
func Remove(users []user.User) []user.User {
	// TODO: you can use struct{} instead of bool, as it occupies 0 space
	allKeys := make(map[user.User]struct{})
	var list []user.User
	for _, item := range users {
		// TODO: a more clear idiom here is _,ok:= or _, found :=
		if _, ok := allKeys[item]; !ok {
			allKeys[item] = struct{}{}
			list = append(list, item)
		}
	}
	return list
}
