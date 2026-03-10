package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

// the put method is safer to user when creating or updating the data when it comes to the updating. 

func updateUser(baseURL, id, apiKey string, data User) (User, error) {
	fullURL := baseURL + "/" + id 

	//marshaling the received data.
	jsonData, err := json.Marshal(data)
	if err != nil {
		return User{}, err
	}

	//creating the new http request.
	req, err := http.NewRequest("PUT", fullURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return User{}, err
	}

	//setting the headers 
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-key", apiKey)


	// making an http request, you can also use http.Default
	client := &http.Client{
		Timeout: 10 * time.Second,
	}  // more efficient with the 10 sec timeout.
	res, err := client.Do(req)
	if err != nil {
		return User{}, err
	}
	defer req.Body.Close()


	var user User
	//decoding the request body and returning it.
	if err := json.NewDecoder(res.Body).Decode(&user); err != nil {
		return user, err// here the user is declared so can't use user{}
	}

	return user, nil
}

func getUserById(baseURL, id, apiKey string) (User, error) {

	fullURL := baseURL + "/" + id

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return User{}, err
	}

	req.Header.Set("X-API-key", apiKey)

	res, err := http.DefaultClient.Do(req) // this DefaultClient waits indefinately. So use &http.Client{} with timeout
	if err != nil {
		return User{}, err
	}

	defer res.Body.Close()

	var user User
	if err := json.NewDecoder(res.Body).Decode(&user); err != nil {
		return user, err
	}
	return user, nil
}
