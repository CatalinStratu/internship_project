package service

import (
	"awesomeProject2/user"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type IRead interface {
	readDates() ([]byte, error)
}

type Input struct {
	Link     string
	Elements int
}

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

func ReadUsers(readDates IRead, elements int) ([]user.User, error) {
	var users []user.User
	for {
		var tempResult []user.User
		body, err := readDates.readDates()

		if err != nil {
			return nil, fmt.Errorf("JSON ERROR: %v", err)
		}

		if err := json.Unmarshal(body, &tempResult); err != nil { // Parse []byte to the go struct pointer
			return nil, fmt.Errorf("can not unmarshal JSON: %v", err)
		}

		users = append(users, tempResult...)
		if len(users) >= elements {
			users = users[:elements]
			break
		}
	}
	return users, nil
}
