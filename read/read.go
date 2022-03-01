package read

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"service/user"
)

// Read interface
type Read interface {
	readDates(client HttpClient, ctx context.Context) ([]byte, error)
}

// Input structure
type Input struct {
	Link     string
	Elements int
}

func (r *Input) readDates(client HttpClient, ctx context.Context) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, r.Link, nil)
	if err != nil {
		return nil, fmt.Errorf("cannot create request: %w", err)
	}

	rsp, ok := client.Do(req)

	if ok != nil {
		return nil, fmt.Errorf("invalide URL")
	}

	defer rsp.Body.Close()

	if rsp.StatusCode > 299 {
		return nil, fmt.Errorf("response failed with status code: %v", rsp.StatusCode)
	}

	body, ok := io.ReadAll(rsp.Body)
	if ok != nil {
		return nil, fmt.Errorf("cannot read all response body: %v", err)
	}

	return body, nil
}

// Users read users from the source return list of users
// list of users is limited by the number of elements
func Users(readDates Read, elements int, ctx context.Context) ([]user.User, error) {
	var users []user.User
	client := sendRequest{}
	for {
		// temporal element
		var tempData []user.User

		// read dates from source
		body, err := readDates.readDates(client, ctx)

		if err != nil {
			return nil, fmt.Errorf("%w", err)
		}

		// Parse []byte to the go struct pointer
		if err = json.Unmarshal(body, &tempData); err != nil {
			return nil, fmt.Errorf("can not unmarshal JSON: %w", err)
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
