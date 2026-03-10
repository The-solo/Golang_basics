package main

// There is no while loop in golang

func getMaxMessagesToSend(costMultiplier float64, maxCostInPennies int) int {
	actualCostInPennies := 1.0
	maxMessagesToSend := 1
	balance := float64(maxCostInPennies) - actualCostInPennies
	// There is no while loop we just use the for loop with the condition.
	for balance >= 0.0 {
		actualCostInPennies *= costMultiplier
		balance -= actualCostInPennies
		maxMessagesToSend++
	}
	// And we specify the breaking condition to avoid the infinite loop
	if balance < 0 {
		maxMessagesToSend--
	}
	return maxMessagesToSend
}
