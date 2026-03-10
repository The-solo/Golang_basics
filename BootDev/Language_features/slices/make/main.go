package main

func getMessageCosts(messages []string) []float64 {
	costs := make([]float64, len(messages))// pre-allocating memory for slice of length messages
	for i:=0; i<len(messages); i++{
		costs[i] =  float64(len(messages[i])) * 0.01
	}
	return costs
}

