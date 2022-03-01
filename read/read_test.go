package read

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"service/user"
	"testing"
)

type MockInputError struct {
	Link     string
	Elements int
}

func (m MockInputError) readDates(client HttpClient, ctx context.Context) ([]byte, error) {
	return nil, fmt.Errorf("cannot create request")
}

type sendRequestError struct{}

func (s sendRequestError) Do(req *http.Request) (*http.Response, error) {
	r := &http.Response{
		Status:     "500",
		StatusCode: 500,
		Body:       ioutil.NopCloser(bytes.NewReader(nil)),
	}

	return r, nil
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

func (m MockInputSuccess) readDates(client HttpClient, ctx context.Context) ([]byte, error) {
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

func TestReadDatesErrorInvalidLink(t *testing.T) {
	ctx := context.Background()
	var input = &Input{Elements: 3, Link: "https://example.coms1234123123"}
	client := sendRequest{}
	_, err := input.readDates(client, ctx)
	if err == nil {
		t.Errorf("invalide URL")
	}
}

func TestReadDatesErrorStatusCodeError(t *testing.T) {
	ctx := context.Background()
	var input = &Input{Elements: 3, Link: "#"}
	client := sendRequestError{}
	_, err := input.readDates(client, ctx)
	if err == nil {
		t.Errorf("Status code error: %v", err)
	}
}
