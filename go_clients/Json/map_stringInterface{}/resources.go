package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
)

func getResources(url string) ([]map[string]any, error) { // any is alias for interface{}
	var resources []map[string]any

	res, err := http.Get(url)
	if err != nil {
		return resources, err
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body) //decoding the input stream.
	if err := decoder.Decode(&resources); err != nil{ // assigning the data to the structure of resources.
		return resources, err
	}
	return resources, nil
}

func logResources(resources []map[string]any) {
	var formattedStrings []string

	for _, data := range resources{
		for key, value := range data{
			formattedStrings = append(formattedStrings, fmt.Sprintf("key:%s - value:%v", key, value))
		}
	}
	sort.Strings(formattedStrings)

	for _, str := range formattedStrings {
		fmt.Println(str)
	}
}

