package main

import ("unicode/utf8")
func getNameCounts(names []string) map[rune]map[string]int {

    nestedMap := make(map[rune]map[string]int)
    for _, name := range names {
        if len(name) == 0 {
            continue // Skip empty strings to avoid errors
        }
        //char := []rune(name)[0] // We convert string to slice of []rune and grab the index 0th element.
		char, _ := utf8.DecodeRuneInString(name) // this func returns 2 values.
        if _, ok := nestedMap[char]; !ok {
            nestedMap[char] = make(map[string]int) 
        }
		// 3. The Critical Step: Initialize the inner map if it doesn't exist
        // If we don't do this, we are trying to write to a 'nil' map
		nestedMap[char][name]++  // Turning the new value immdietly to 1. AKA counter
		//We can use ths because inner map has value of type int.
    }
    return nestedMap
}



/*
//
//The alternate method using struct to keep the track of the counter for each name occurance. 


type NameRecord struct {
    Count int
    Exists bool // This is your "bool flag"
}

func getNameCounts(names []string) map[rune]map[string]NameRecord {
    nestedMap := make(map[rune]map[string]NameRecord)

    for _, name := range names {
        if len(name) == 0 { continue }
       // char := []rune(name)[0]
		char, _ := utf8.DecodeRuneInString(name) // this func returns 2 values.

        // Ensure the "drawer" exists (same as before)
        if _, ok := nestedMap[char]; !ok {
            nestedMap[char] = make(map[string]NameRecord)
        }

        // Get the current record (C-brain: fetching the struct)
        record := nestedMap[char][name]

        // Update the record
        record.Count++
        record.Exists = true // Mark your bool flag

        // Put it back in the map
        nestedMap[char][name] = record
    }
    return nestedMap
}

*/
