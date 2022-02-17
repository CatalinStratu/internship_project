package service

import (
	"awesomeProject2/user"
	"fmt"
	"reflect"
	"testing"
)

type MockIReadError interface {
	readDates() ([]byte, error)
}

type MockInputError struct {
	Link     string
	Elements int
}

func (m MockInputError) readDates() ([]byte, error) {
	return nil, fmt.Errorf("cannot create request")
}

func TestReadUsersReadDatesError(t *testing.T) {
	var input = &MockInputError{Elements: 100, Link: "https://randomapi.com/api/6de6abfedb24f889e0b5f675edc50deb?fmt=raw&sole"}
	removeDuplicate, _ := ReadUsers(input, input.Elements)
	var expected []user.User

	if !reflect.DeepEqual(removeDuplicate, expected) {
		t.Errorf("read dates error, got: %v, want: %v.", expected, removeDuplicate)
	}
}

type MockIReadSuccess interface {
	readDates() ([]byte, error)
}

type MockInputSuccess struct {
	Link     string
	Elements int
}

func (m MockInputSuccess) readDates() ([]byte, error) {
	var s []byte
	return s, nil
}

func TestReadUsersReadDatesSuccess(t *testing.T) {
	var input = &MockInputSuccess{Elements: 100, Link: "https://randomapi.com/api/6de6abfedb24f889e0b5f675edc50deb?fmt=raw&sole"}
	removeDuplicate, _ := ReadUsers(input, input.Elements)
	var expected []user.User

	if !reflect.DeepEqual(removeDuplicate, expected) {
		t.Errorf("Read dates succes, got: %v, want: %v.", expected, removeDuplicate)
	}
}

type MockIReadReadDatesTest interface {
	readDates() ([]byte, error)
}

type MockInputReadDatesTest struct {
	Link     string
	Elements int
}

func TestReadDatesErrorStatusCode(t *testing.T) {
	var input = &Input{Elements: 3, Link: "https://randomapi.com/api/16de6abfedb24f889e0b5f675edc50deb?fmt=raw&sole"}
	_, err := input.readDates()
	if err == nil {
		t.Errorf("Status code error: %v", err)
	}
}

func TestReadDatesErrorInvalidLink(t *testing.T) {
	var input = &Input{Elements: 3, Link: "https://randomapi.coms1234123123"}
	_, err := input.readDates()
	if err == nil {
		t.Errorf("invalide URL")
	}
}
