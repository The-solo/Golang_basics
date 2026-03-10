package main

func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {
	messages :=[3]string{primary, secondary, tertiary}
	first := len(primary)
	second := len(secondary) + first
	third := len(tertiary) + second
	lengths := [3]int{first, second, third}
	return messages, lengths
}

