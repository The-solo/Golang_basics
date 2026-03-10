package main

import (
	"errors"
)

const (
	planFree = "free"
	planPro  = "pro"
)

func getMessageWithRetriesForPlan(plan string, messages [3]string) ([]string, error) {
	if plan == "pro" {
		output := messages[:]
		return output, nil
	} else if plan == "free" {
		output := messages[0:2]
		return output, nil
	} else {
		var output []string 
		return output,  errors.New("unsupported plan")
	}
}

