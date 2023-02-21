package cmd

import (
	"net/http"
	"time"
)


type AudiofileClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var (
	getClient = GetHTTPClient()
) 

func GetHTTPClient() AudiofileClient {
	return &http.Client{
		Timeout: 15 * time.Second,
	}
}
