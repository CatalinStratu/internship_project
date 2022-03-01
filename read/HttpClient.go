package read

import "net/http"

//HttpClient interface
type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type sendRequest struct{}

//Do send an HTTP request and returns an HTTP response
func (s sendRequest) Do(req *http.Request) (*http.Response, error) {
	client := http.Client{}
	return client.Do(req)
}
