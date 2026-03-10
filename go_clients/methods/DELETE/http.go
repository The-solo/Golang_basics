package main

import (
	"fmt"
	"net/http"
	"time"
)

func deleteUser(baseURL, id, apiKey string) error {
	fullURL := baseURL + "/" + id

	req, err := http.NewRequest("DELETE", fullURL, nil)
	if err != nil {
		 return err
	}

	req.Header.Set("X-API-key", apiKey)

	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode > 299 {
		fmt.Println("Request to delete location unsuccessfull!")
	}
	return nil
}
