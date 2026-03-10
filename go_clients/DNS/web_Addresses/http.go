package main

import (
	"fmt"
	"io"
	"net/http"
	"encoding/json"
)

func getIPAddress(domain string) (string, error) {
	url := fmt.Sprintf("https://1.1.1.1/dns-query?name=%s&type=A", domain)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("accept", "application/dns-json")

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request: %w", err)
	}

	defer res.Body.Close()

	var dnsRes DNSResponse // Declaring the struct locally

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %w", err)
	}

	if err := json.Unmarshal(body, &dnsRes); err != nil { //Reading all data at once.
		return "", err
	}

	if len(dnsRes.Answer) == 0{ // cheaking if empty.
		return "", fmt.Errorf("no answer found")
	}
	//return the first ip address found.
	return dnsRes.Answer[0].Data, nil
}

