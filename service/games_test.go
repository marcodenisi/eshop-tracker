package service

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestRetrieveGames(t *testing.T) {
	resp, err := RetrieveEuGames(&mockClient{})
	if err != nil {
		t.Errorf("Got err %v instead of nil", err)
	}

	if len(resp) != 2 {
		t.Errorf("Got %v games instead of 2", len(resp))
	}
}

type mockClient struct{}

func (mc *mockClient) Do(req *http.Request) (*http.Response, error) {
	mock, _ := ioutil.ReadFile("../stubs/eugames.json")
	r := ioutil.NopCloser(bytes.NewReader(mock))

	return &http.Response{
		Status:     string(http.StatusOK),
		StatusCode: http.StatusOK,
		Body:       r,
	}, nil
}
