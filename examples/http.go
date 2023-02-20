package examples

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

func HTTPTimeout() {
	client := http.Client{
		Timeout: 1 * time.Second,
	}
	body := &bytes.Buffer{}
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080/timeout", body)
	if err != nil {
		panic(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		urlErr := err.(*url.Error)
		if urlErr.Timeout() {
			fmt.Println("timeout: ", err)
			return
		}
	}
	defer resp.Body.Close()
}

func HTTPError() {
	resp, err := http.Get("http://localhost:8080/error")
	if err != nil {
		// a non-200 response doesn't cause an error
		// too many redirects
		// any *url.Error type, and the Timeout value will be true
		urlErr := err.(*url.Error)
		if urlErr.Timeout() {
			// a timeout is a type of error
			fmt.Println("timeout: ", err)
			return
		}
		if urlErr.Temporary() {
			// a temporary network error, retry later
			fmt.Println("temporary: ", err)
			return
		}
		fmt.Printf("operation: %s, url: %s, error: %s\n", urlErr.Op, urlErr.URL, urlErr.Error())
		return
	}
	if resp != nil {
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			// action for when status code is not okay
			switch resp.StatusCode {
			case http.StatusBadRequest:
				fmt.Printf("bad request: %v\n", resp.Status)
			case http.StatusInternalServerError:
				fmt.Printf("internal service error: %v\n", resp.Status)
			default:
				fmt.Printf("unexpected status code: %v\n", resp.StatusCode)
			}
		}
		data, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("err:", err)
		}
		fmt.Println("response body:", string(data))
	}
}
