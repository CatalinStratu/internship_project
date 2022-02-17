package grouping

import (
	"awesomeProject2/user"
	"reflect"
	"testing"
)

func TestGrouping(t *testing.T) {
	user1 := user.User{FirstName: "Test Item 1", LastName: "Test", Email: "test", Address: "test", Created: "test", Balance: "test"}
	user2 := user.User{FirstName: "Test Item 1", LastName: "Test", Email: "test", Address: "test", Created: "test", Balance: "test"}
	user3 := user.User{FirstName: "Test Item 1", LastName: "Test", Email: "test", Address: "test", Created: "test", Balance: "test"}
	var users []user.User
	expected := make(map[string][]user.User)
	users = append(users, user1, user2, user3)

	grouping := Grouping(users)

	expected["T"] = append(expected["T"], user1, user2, user3)
	if reflect.DeepEqual(grouping, expected) {
		t.Errorf("Grouping was incorrect, got: %v, want: %v.", expected, expected["T"])
	}
}
