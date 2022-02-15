package main

import (
	"awesomeProject2/duplicate"
	"awesomeProject2/export"
	"awesomeProject2/grouping"
	api "awesomeProject2/service"
)

func main() {
	var users []api.User
	users = api.Users()
	users = duplicate.Remove(users)
	collections := make(map[string][]api.User)
	collections = grouping.Grouping(users)
	export.Export(collections)
}
