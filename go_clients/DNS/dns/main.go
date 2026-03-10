package main

import (
	"net/url"
)

func getDomainNameFromURL(rawURL string) (string, error) {

// Initiating the URL struct blueprint using url.Parse
	parsedURL, err := url.Parse(rawURL)
	if err != nil{
		return "", err
	}
	//Returning the extracted hostname with this function call
	return parsedURL.Hostname(), nil

}

