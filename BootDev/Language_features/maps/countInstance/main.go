package main

func updateCounts(messagedUsers []string, validUsers map[string]int) {
	for i := range messagedUsers {
		if _, ok := validUsers[messagedUsers[i]]; ok{
			validUsers[messagedUsers[i]]++
		}
	}
}
