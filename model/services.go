package model

import "net/http"

// HTTPClient is an interface to allow http client mocking
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
