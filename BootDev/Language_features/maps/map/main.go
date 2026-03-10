package main

import "errors"

//Maps are object stores where the data is in the form of key value pairs

func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	phonebook := make(map[string]user)
	if len(names) != len(phoneNumbers) {
		return phonebook, errors.New("invalid sizes")
	}
	for i, name := range names{ //since both slices are of same size the phoneNumbers should work as well.
		phonebook[name]= user {
			name : name,
			phoneNumber : phoneNumbers[i],
		} // This is how you update the map with the dynamic values of struct.
	}
	return phonebook, nil
}

type user struct {
	name        string
	phoneNumber int
}
