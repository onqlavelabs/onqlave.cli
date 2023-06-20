package cli

import (
	"net/http"
)

func CheckResponseError(r *http.Response) error {
	if r.StatusCode >= 200 && r.StatusCode < 300 {
		return nil
	}

	return nil
}
