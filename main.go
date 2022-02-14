package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Response struct {
	FirstName string `json:"first"`
	LastName  string `json:"last"`
	Email     string `json:"email"`
	Address   string `json:"address"`
	Created   string `json:"created"`
	Balance   string `json:"balance"`
}

type Records struct {
	Index   string     `json:"index"`
	Records []Response `json:"records"`
	Total   int        `json:"total_records"`
}

func main() {
	var result []Response

	for i := 0; i < 1; i++ {
		resp, err := http.Get("https://randomapi.com/api/6de6abfedb24f889e0b5f675edc50deb?fmt=raw&sole")
		if err != nil {
			fmt.Println("No response from request")
		}
		var tempResult []Response
		body, err := ioutil.ReadAll(resp.Body)                    // response body is []byte
		if err := json.Unmarshal(body, &tempResult); err != nil { // Parse []byte to the go struct pointer
			fmt.Println("Can not unmarshal JSON")
		}
		resp.Body.Close()
		result = append(result, tempResult...)
	}

	item1 := Response{FirstName: "Test Item 1", LastName: "Test", Email: "test", Address: "test", Created: "test", Balance: "test"}
	item2 := Response{FirstName: "Test Item 1", LastName: "Test", Email: "test", Address: "test", Created: "test", Balance: "test"}
	item3 := Response{FirstName: "Test Item 1", LastName: "Test", Email: "test", Address: "test", Created: "test", Balance: "test"}

	result = append(result, item1, item2, item3)

	result = removeDuplicateResponses(result)

	collections := make(map[string][]Response)
	for _, b := range result {
		collections[b.FirstName[0:1]] = append(collections[b.FirstName[0:1]], b)
	}

	var records []Records

	for i, j := range collections {
		tempRecord := Records{Index: i, Records: j, Total: len(j)}
		records = append(records, tempRecord)
		file, _ := json.MarshalIndent(tempRecord, "", " ")
		var name string
		name = i + ".json"
		_ = ioutil.WriteFile(name, file, 0644)
	}
}

func removeDuplicateResponses(resSlice []Response) []Response {
	allKeys := make(map[Response]bool)
	var list []Response
	for _, item := range resSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}
