package main

import (
	"awesomeProject2/duplicate"
	"awesomeProject2/export"
	"awesomeProject2/grouping"
	api "awesomeProject2/service"
	"awesomeProject2/user"
	"fmt"
)

func main() {
	// User input
	var input = &api.Input{Elements: 50, Link: "https://randomapi.com/api/6de6abfedb24f889e0b5f675edc50deb?fmt=raw&sole"}
	users, err := api.ReadUsers(input, input.Elements)
	if err != nil {
		fmt.Printf("Read users error: %v", err)
	}
	users = duplicate.Remove(users)
	// Create collections for grouping
	collections := make(map[string][]user.User)
	collections = grouping.Grouping(users)
	var w = &export.Write{}
	err = export.Export(collections, w)
	if err != nil {
		fmt.Printf("Export error: %v", err)
	}
}
