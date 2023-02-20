package cmd

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

// mock client

type ClientMock struct {
}

func (c *ClientMock) Do(req *http.Request) (*http.Response, error) {
	listBytes, err := os.ReadFile("./testfiles/list.json")
	if err != nil {
		return nil, fmt.Errorf("unable to read testfile/list.json")
	}
	getBytes, err := os.ReadFile("./testfiles/get.json")
	if err != nil {
		return nil, fmt.Errorf("unable to read testfile/get.json")
	}
	searchBytes, err := os.ReadFile("./testfiles/search.json")
	if err != nil {
		return nil, fmt.Errorf("unable to read testfile/search.json")
	}
	switch req.URL.Path {
	case "/request":
		return &http.Response{
			Status:        "OK",
			StatusCode:    http.StatusOK,
			Body:          io.NopCloser(bytes.NewBufferString(string(getBytes))),
			ContentLength: int64(len(getBytes)),
			Request:       req,
			Header:        make(http.Header, 0),
		}, nil
	case "/upload":
		return &http.Response{
			Status:        "OK",
			StatusCode:    http.StatusOK,
			Body:          io.NopCloser(bytes.NewBufferString("123")),
			ContentLength: int64(len("123")),
			Request:       req,
			Header:        make(http.Header, 0),
		}, nil
	case "/list":
		return &http.Response{
			Status:        "OK",
			StatusCode:    http.StatusOK,
			Body:          io.NopCloser(bytes.NewBufferString(string(listBytes))),
			ContentLength: int64(len(listBytes)),
			Request:       req,
			Header:        make(http.Header, 0),
		}, nil
	case "/delete":
		return &http.Response{
			Status:        "OK",
			StatusCode:    http.StatusOK,
			Body:          io.NopCloser(bytes.NewBufferString("successfully deleted audio with id: 456")),
			ContentLength: int64(len("successfully deleted audio with id: 456")),
			Request:       req,
			Header:        make(http.Header, 0),
		}, nil
	case "/search":
		return &http.Response{
			Status:        "OK",
			StatusCode:    http.StatusOK,
			Body:          io.NopCloser(bytes.NewBufferString(string(searchBytes))),
			ContentLength: int64(len(searchBytes)),
			Request:       req,
			Header:        make(http.Header, 0),
		}, nil
	}
	return &http.Response{}, nil
}
