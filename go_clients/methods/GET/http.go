package main

import (
	"encoding/json"
	"net/http"
	"fmt"
)

func getUsers(url string) ([]User, error) {

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Response body is empty", err)
	}
	defer resp.Body.Close()

	var user []User
	//decoding the response data and mapping it to the User[] struct
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&user); err != nil {
		return nil, fmt.Errorf("Something went wrong!", err)
	}
	return user, nil
}
