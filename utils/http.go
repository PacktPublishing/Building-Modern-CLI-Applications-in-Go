package utils

import (
	"fmt"
	"net/http"
)

func CheckResponse(resp *http.Response) error {
	if resp != nil {
		if resp.StatusCode != http.StatusOK {
			// action for when status code is not okay
			return fmt.Errorf("unexpected response: %v", resp.Status)
		}
		return nil
	} else {
		return fmt.Errorf("response body is nil")
	}
}
