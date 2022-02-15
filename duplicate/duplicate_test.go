package duplicate

import (
	api "awesomeProject2/service"
	"reflect"
	"testing"
)

func TestRemove(t *testing.T) {
	user1 := api.User{FirstName: "Test Item 1", LastName: "Test", Email: "test", Address: "test", Created: "test", Balance: "test"}
	user2 := api.User{FirstName: "Test Item 1", LastName: "Test", Email: "test", Address: "test", Created: "test", Balance: "test"}
	user3 := api.User{FirstName: "Test Item 1", LastName: "Test", Email: "test", Address: "test", Created: "test", Balance: "test"}
	var users []api.User
	users = append(users, user1, user2, user3)

	removeDuplicate := Remove(users)

	var expected []api.User

	expected = append(expected, user1)

	if !reflect.DeepEqual(removeDuplicate, expected) {
		t.Errorf("Array chunk was incorrect, got: %v, want: %v.", expected, removeDuplicate)
	}
}
