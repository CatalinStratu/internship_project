package service

import (
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

func Users() []User {
	var users []User
	for i := 0; i < 1; i++ {
		resp, err := http.Get("https://randomapi.com/api/6de6abfedb24f889e0b5f675edc50deb?fmt=raw&sole")
		if err != nil {
			fmt.Println("No response from request")
		}
		var tempResult []User
		body, err := io.ReadAll(resp.Body)                        // response body is []byte
		if err := json.Unmarshal(body, &tempResult); err != nil { // Parse []byte to the go struct pointer
			fmt.Println("Can not unmarshal JSON")
		}
		resp.Body.Close()
		users = append(users, tempResult...)
	}
	return users
}
