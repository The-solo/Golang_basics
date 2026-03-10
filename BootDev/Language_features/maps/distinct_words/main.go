package main

import(
	"strings"
)
/*

// my Origial implementation.

func countDistinctWords(messages []string) int {
	counter := 0
	records := make(map[string]bool)
	for _, msg:= range messages{
		new_word := strings.Fields(strings.ToLower(msg)) // returns the slice of string values
		for _, unique := range new_word {
			if ok := records[unique]; ok{
				continue
			}
			records[unique] = true
			counter++
		}
	}
	return counter
}
*/


// The struct method is more efficient because struct here only track presence.
func countDistinctWords(messages []string) int {

    records := make(map[string]struct{})
    for _, msg := range messages {
        words := strings.Fields(strings.ToLower(msg))
        for _, word := range words {

            records[word] = struct{}{}  // The empty struct here has a value of 0 bytes.
		// The 2nd bracket is the value of type struct{}, because the struct{} is a type not a values
		// This satisfies the map that every key has a value (struct{} key has a value {} as in nothing)
        }
    }
    return len(records) // len returns the count for keys present.
}

