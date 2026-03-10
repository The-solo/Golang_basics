package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)

func getIssues(url string) ([]Issue, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	defer res.Body.Close()

	var issues []Issue // creating the slice of type Issue
	decoder := json.NewDecoder(res.Body)// creating a new decode

	if err := decoder.Decode(&issues); err != nil{ // .
		return nil, fmt.Errorf("error decoding the json : %w", err)
	}
	return issues, nil
}
