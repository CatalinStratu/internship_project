package read

import (
	"context"
	"fmt"
	"reflect"
	"service/user"
	"testing"
)

// TODO: not a good idea to test with real data!!!!
// TODO: you don't need to redefine the interface!!!
type MockInputError struct {
	Link     string
	Elements int
}

func (m MockInputError) readDates(ctx context.Context) ([]byte, error) {
	return nil, fmt.Errorf("cannot create request")
}

func TestReadUsersReadDatesError(t *testing.T) {
	ctx := context.Background()
	var input = &MockInputError{Elements: 100, Link: ""}
	removeDuplicate, _ := Users(input, input.Elements, ctx)
	var expected []user.User

	if !reflect.DeepEqual(removeDuplicate, expected) {
		t.Errorf("read dates error, got: %v, want: %v.", expected, removeDuplicate)
	}
}

type MockInputSuccess struct {
	Link     string
	Elements int
}

func (m MockInputSuccess) readDates(ctx context.Context) ([]byte, error) {
	var s []byte
	return s, nil
}

func TestReadUsersReadDatesSuccess(t *testing.T) {
	ctx := context.Background()
	var input = &MockInputSuccess{Elements: 100, Link: ""}
	removeDuplicate, _ := Users(input, input.Elements, ctx)
	var expected []user.User

	if !reflect.DeepEqual(removeDuplicate, expected) {
		t.Errorf("Read dates succes, got: %v, want: %v.", expected, removeDuplicate)
	}
}

// Check invalid link
func TestReadDatesErrorInvalidLink(t *testing.T) {
	ctx := context.Background()
	var input = &Input{Elements: 3, Link: "https://example.coms1234123123"}
	_, err := input.readDates(ctx)
	if err == nil {
		t.Errorf("invalide URL")
	}
}

func TestReadDatesErrorStatusCode(t *testing.T) {
	ctx := context.Background()
	var input = &Input{Elements: 3, Link: "https://randomapi.com/api/16de6abfedb24f889e0b5f675edc50deb?fmt=raw&sole"}
	_, err := input.readDates(ctx)
	if err != nil {
		t.Errorf("Status code error: %v", err)
	}
}
