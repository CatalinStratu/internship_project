package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type User struct {
	FirstName string `json:"first"`
	LastName  string `json:"last"`
	Email     string `json:"email"`
	Address   string `json:"address"`
	Created   string `json:"created"`
	Balance   string `json:"balance"`
}

type Service interface {
	Users()
}

func Request() ([]byte, error) {
	client := http.Client{}
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "https://randomapi.com/api/6de6abfedb24f889e0b5f675edc50deb?fmt=raw&sole", nil)
	if err != nil {
		return nil, fmt.Errorf("cannot create request: %v", err)
	}

	rsp, err := client.Do(req)
	if rsp != nil {
		defer rsp.Body.Close()
	}

	if rsp.StatusCode == 404 || rsp.StatusCode == 403 || rsp.StatusCode == 401 || rsp.StatusCode == 500 {
		return nil, fmt.Errorf("response failed with status code: %v", rsp.StatusCode)
	}

	body, err := io.ReadAll(rsp.Body)
	if err != nil {
		return nil, fmt.Errorf("cannot read all response body: %v", err)
	}
	return body, nil
}

func Users() ([]User, error) {
	var users []User
	for i := 0; i < 5; i++ {
		var tempResult []User
		body, err := Request()
		if err != nil {
			return nil, fmt.Errorf("JSON ERROR: %v", err)
		}
		if err := json.Unmarshal(body, &tempResult); err != nil { // Parse []byte to the go struct pointer
			return nil, fmt.Errorf("can not unmarshal JSON: %v", err)
		}

		users = append(users, tempResult...)
	}
	return users, nil
}
