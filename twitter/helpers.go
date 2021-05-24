package twitter

import (
	"fmt"
	"net/url"
)

// appendQueryParams each key/value pair will be appended to the provided endpoint as URL query parameters
//
// The new URL will be returned as a string value
func appendQueryParams(endpoint string, queryParams map[string]string) (string, error) {
	u, err := url.Parse(endpoint)
	if err != nil {
		return "", fmt.Errorf("failed to parse URL; %w", err)
	}

	q := u.Query()

	for k, v := range queryParams {
		q.Set(k, v)
	}

	u.RawQuery = q.Encode()

	return u.String(), nil
}
