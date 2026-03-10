package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func createUser(url, apiKey string, data User) (User, error) {

	//Encoding the data as the Json 
	jsonData, err := json.Marshal(data)
	if err != nil {
		return User{}, err
	}

	//Creating a new post requesta 			the bytes.NewBuffer implements the io.Reader
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData)) //jsonData is the byte[]
	if err != nil {
		return User{}, err
	}

	//setting the headers.
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-key", apiKey)

	//making an http request & storing the response inside 'res'
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return User{}, err
	}
	defer res.Body.Close()

	// Decodind the response data from the response
	var user User
	if err := json.NewDecoder(res.Body).Decode(&user); err != nil { //reading from the stream instead of bytes.
		return user, err
	}
	return user, nil
}
