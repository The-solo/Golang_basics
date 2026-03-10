package main

import (
	"fmt"
)

func validateStatus(status string) error {
	length := len(status)
	if length == 0 {
		return fmt.Errorf("status cannot be empty")
	}
	if length > 140 {
		return fmt.Errorf("status exceeds 140 characters")
	}	
	return nil
}

