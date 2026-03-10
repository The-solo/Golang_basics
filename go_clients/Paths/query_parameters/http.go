package main

import (
	"encoding/json"
	"net/http"
)

func getUsers(url string) ([]User, error) {

	fullURL := url+"?sort=experience" // the query parameter for sorting users by experience.
	res, err := http.Get(fullURL)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var users []User

	if err := json.NewDecoder(res.Body).Decode(&users); err != nil {
		return nil, err
	}
	
	return users, nil
}
