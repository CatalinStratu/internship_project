package main

import (
	"awesomeProject2/duplicate"
	"awesomeProject2/export"
	"awesomeProject2/grouping"
	api "awesomeProject2/service"
	"awesomeProject2/user"
)

func main() {
	var users []user.User
	var input = &api.Input{Elements: 3, Link: "https://randomapi.com/api/6de6abfedb24f889e0b5f675edc50deb?fmt=raw&sole"}
	users, _ = api.ReadUsers(input, input.Elements)
	users = duplicate.Remove(users)
	collections := make(map[string][]user.User)
	collections = grouping.Grouping(users)
	err := export.Export(collections)
	if err != nil {
		return
	}
}
