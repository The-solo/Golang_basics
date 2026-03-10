package main

import (
	"strconv"
	"strings"
)

func fetchTasks(baseURL, availability string) []Issue {

	var limit int

	switch strings.ToLower(availability) {
	case "low":
		limit = 1
	case "medium":
		limit = 3
	case "high":
		limit = 5
	}

	fullURL := baseURL + "?sort=estimate&limit=" + strconv.Itoa(limit)
	return getIssues(fullURL)
}
