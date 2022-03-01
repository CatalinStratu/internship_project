package grouping

import (
	"reflect"
	"testing"

	"service/user"
)

var user1 = user.User{FirstName: "Test Item 1", LastName: "Test", Email: "test", Address: "test", Created: "test", Balance: "test"}
var user2 = user.User{FirstName: "Test Item 1", LastName: "Test", Email: "test", Address: "test", Created: "test", Balance: "test"}
var user3 = user.User{FirstName: "", LastName: "Test", Email: "test", Address: "test", Created: "test", Balance: "test"}

// TODO: failing test
// TODO: add a test with multiple letters
func TestGrouping(t *testing.T) {
	var users []user.User
	expected := make(map[string][]user.User)
	users = append(users, user1, user2, user3)

	grouping := Execute(users)

	expected["T"] = append(expected["T"], user1, user2, user3)
	if reflect.DeepEqual(grouping, expected) {
		t.Errorf("Grouping was incorrect, got: %v, want: %v.", expected, expected["T"])
	}
}
