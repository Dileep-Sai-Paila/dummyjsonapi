// File: pkg/utils/http.go

package utils

import (
	"fmt"
	"net/http"
	"time"
)

// to use a configurable HTTP client with a timeout, its like creating a single client instance and re-using it rather than creating a new one for every request.
var httpClient = &http.Client{
	Timeout: 10 * time.Second,
}

// utility function that performs HTTP GET request to the soecified URL using my own httpClient
func MakeGETRequest(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil) // creating a new request object.
	if err != nil {
		return nil, fmt.Errorf("failed to create request object: %w", err)
	}

	// Perform the request using our shared, timeout-configured client.
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request to %s: %w", url, err)
	}

	return resp, nil
}
