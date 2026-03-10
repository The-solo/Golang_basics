package main

func indexOfFirstBadWord(msg []string, badWords []string) int {
	for i, message := range msg{
		for _, word:= range badWords{
			if word == message{
				return i 
			}
		}
	} 
	return -1
}


