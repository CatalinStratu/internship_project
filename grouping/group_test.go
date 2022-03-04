package grouping

import (
	"reflect"
	"service/user"
	"testing"
)

var user1 = user.User{FirstName: "Test Item 1", LastName: "Test", Email: "test", Address: "test", Created: "test", Balance: "test"}
var user2 = user.User{FirstName: "Test Item 1", LastName: "Test", Email: "test", Address: "test", Created: "test", Balance: "test"}
var user3 = user.User{FirstName: "", LastName: "Test", Email: "test", Address: "test", Created: "test", Balance: "test"}
var user4 = user.User{FirstName: "Keith", LastName: "Hilpert", Email: "Keith.Hilpert@sammy.name", Address: "230 Kenna Harbor", Created: "March 11, 2019", Balance: "$8,419.47"}
var user5 = user.User{FirstName: "Brendan", LastName: "Mosciski", Email: "Keith.Hilpert@sammy.name", Address: "230 Kenna Harbor", Created: "March 11, 2019", Balance: "$8,419.47"}

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

func TestGroupingMultipleLetters(t *testing.T) {
	var users []user.User
	expected := make(map[string][]user.User)
	users = append(users, user1, user2, user3, user4, user5)

	grouping := Execute(users)

	expected["T"] = append(expected["T"], user1, user2, user3)
	expected["B"] = append(expected["B"], user5)
	expected["K"] = append(expected["K"], user4)

	if reflect.DeepEqual(grouping, expected) {
		t.Errorf("Grouping was incorrect, got: %v, want: %v.", expected, expected)
	}
}
