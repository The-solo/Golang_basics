package main

import (
	"encoding/json"
)

// The go struct values are assemebled into bytes/json to be sent.

func marshalAll[T any](items []T) ([][]byte, error) {
	var final [][]byte
	for _, item := range items{
		data, err := json.Marshal(item) // converting the JSON data into bytes.
		if err != nil{
			return nil, err
		}
		final = append(final, data)
	}
	return final, nil
}

