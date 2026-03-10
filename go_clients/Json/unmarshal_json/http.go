package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// The incoming data can be raw bytes/json will be converted to go values so go code can access it.

func getIssues(url string) ([]Issue, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body) // reading data all at once
	if err != nil {
		return nil, err
	}

	var issues []Issue
	if err := json.Unmarshal(data, &issues); err != nil { //converts the raw bytes to json values.
		return nil, err
	}
	return issues, nil
}
