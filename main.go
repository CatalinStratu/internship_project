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
	var input = &read.Input{Elements: 1000, Link: "https://randomapi.com/api/6de6abfedb24f889e0b5f675edc50deb?fmt=raw&sole"}
	users, err := read.Users(input, input.Elements, ctx)
	// TODO: your app should not continue if you can't read data!!!
	if err != nil {
		fmt.Printf("Read users error: %v", err)
		return
	}
	users = duplicate.Remove(users)
	// TODO: why are you creating the map in collections if you  overwrite it right away?
	// TODO: grouping.Grouping - nice name!!!
	collections := grouping.Execute(users)
	var w = &export.WriteFile{}
	// TODO: export.Export sounds very "nice" :)
	err = export.Execute(collections, w)
	if err != nil {
		fmt.Printf("Export error: %v", err)
		return
	}
}
