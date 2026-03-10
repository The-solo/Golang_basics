package main

import (
	"strings"
)

func removeProfanity(message *string) { //message is a pointer to data of type string

// dereferencing the pointer message with *message
	if strings.Contains(*message, "fubb") {
		*message = strings.ReplaceAll(*message, "fubb", "****")
	}
	if strings.Contains(*message, "shiz") {
		*message = strings.ReplaceAll(*message, "shiz", "****")
	}
	*message = strings.ReplaceAll(*message, "witch", "*****")
}
