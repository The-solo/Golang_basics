package main

import (
	"fmt"
	"io"
	"net/http"
)

func getIssueData(url string) ([]byte, error) {
	res, err := http.Get(url) // sends the get request to the url. uses http.DefaultClient
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body) //Reads the data into the "Slice of bytes"
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	return data, nil
}
