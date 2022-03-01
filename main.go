package main

import (
	"context"
	"fmt"
	"service/duplicate"
	"service/export"
	"service/grouping"
	// TODO: don't use hacks like import alias without a good reason(like a conflicting name)
	"service/read"
)

func main() {
	ctx := context.Background()
	// User input
	var input = &read.Input{Elements: 100, Link: "https://randomapi.com/api/6de6abfedb24f889e0b5f675edc50deb?fmt=raw&sole"}
	users, err := read.Users(input, input.Elements, ctx)
	if err != nil {
		fmt.Printf("Read users error: %v", err)
		return
	}
	users = duplicate.Remove(users)
	collections := grouping.Execute(users)
	var w = &export.WriteFile{}
	err = export.Execute(collections, w)
	if err != nil {
		fmt.Printf("Export error: %v", err)
		return
	}
}
