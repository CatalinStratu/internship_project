package service

import (
	"awesomeProject2/user"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// IRead interface
type IRead interface {
	readDates() ([]byte, error)
}

// Input structure
type Input struct {
	Link     string
	Elements int
}

// readDates Read data from rest api
func (r *Input) readDates() ([]byte, error) {
	client := http.Client{}
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, r.Link, nil)
	if err != nil {
		return nil, fmt.Errorf("cannot create request: %v", err)
	}

	rsp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("invalide URL")

	}
	if rsp != nil {
		defer rsp.Body.Close()
	}

	if rsp.StatusCode > 299 {
		return nil, fmt.Errorf("response failed with status code: %v", rsp.StatusCode)
	}

	body, err := io.ReadAll(rsp.Body)
	if err != nil {
		return nil, fmt.Errorf("cannot read all response body: %v", err)
	}
	return body, nil
}

// ReadUsers read users from the source return list of users
// list of users is limited by the number of elements
func ReadUsers(readDates IRead, elements int) ([]user.User, error) {
	var users []user.User
	for {
		// temporal element
		var tempData []user.User

		// read dates from source
		body, err := readDates.readDates()

		if err != nil {
			return nil, fmt.Errorf("JSON ERROR: %v", err)
		}

		// Parse []byte to the go struct pointer
		if err := json.Unmarshal(body, &tempData); err != nil {
			return nil, fmt.Errorf("can not unmarshal JSON: %v", err)
		}

		users = append(users, tempData...)
		if len(users) >= elements {
			//delete all items after the "elements" index
			users = users[:elements]
			break
		}
	}
	return users, nil
}
