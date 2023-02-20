package utils

import (
	"fmt"
	"net/http"
)

func CheckResponse(resp *http.Response) error {
	if resp != nil {
		if resp.StatusCode != http.StatusOK {
			switch resp.StatusCode {
			case http.StatusInternalServerError:
				return fmt.Errorf(errorColor("retry the command later"))
			case http.StatusNotFound:
				return fmt.Errorf(errorColor("the id cannot be found"))
			default:
				return fmt.Errorf(errorColor(fmt.Sprintf("unexpected response: %v", resp.Status)))
			}
		}
		return nil
	} else {
		return fmt.Errorf(errorColor("response body is nil"))
	}
}
