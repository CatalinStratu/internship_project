package grouping

import (
	api "awesomeProject2/service"
	"reflect"
	"testing"
)

func TestGrouping(t *testing.T) {
	user1 := api.User{FirstName: "Test Item 1", LastName: "Test", Email: "test", Address: "test", Created: "test", Balance: "test"}
	user2 := api.User{FirstName: "Test Item 1", LastName: "Test", Email: "test", Address: "test", Created: "test", Balance: "test"}
	user3 := api.User{FirstName: "Test Item 1", LastName: "Test", Email: "test", Address: "test", Created: "test", Balance: "test"}
	var users []api.User
	expected := make(map[string][]api.User)
	users = append(users, user1, user2, user3)

	grouping := Grouping(users)

	expected["T"] = append(expected["T"], user1, user2, user3)
	if reflect.DeepEqual(grouping, expected) {
		t.Errorf("Array chunk was incorrect, got: %v, want: %v.", expected, expected["T"])
	}
}
