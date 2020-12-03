package services

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestConsumer_GetUsers(t *testing.T) {

	dat, readError := ioutil.ReadFile("../mocks/open-banking/channels/v1/branches")

	if readError != nil {
		fmt.Println(readError)
	}

	client := NewTestClient(func(req *http.Request) *http.Response {

		return &http.Response{
			StatusCode: 200,
			// Send response to be tested
			Body: ioutil.NopCloser(bytes.NewBufferString(string(dat))),
			// Must be set to non-nil value or it panics
			Header: make(http.Header),
		}
	})

	crawler := NewCrawler(client)

	branches, crawlErr := crawler.Branches("http://teste")
	if crawlErr != nil {
		t.Error(crawlErr)
	}

	if len(*branches) == 0 {
		t.Error("Expected branches")
	}

}

// RoundTripFunc .
type RoundTripFunc func(req *http.Request) *http.Response

// RoundTrip .
func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

//NewTestClient returns *http.Client with Transport replaced to avoid making real calls
func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: RoundTripFunc(fn),
	}
}
