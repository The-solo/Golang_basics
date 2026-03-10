package main

func countConnections(groupSize int) int {
	totalConnections := 0
	for i:=0; i<groupSize; i++{
		totalConnections += i
	}
	return totalConnections
}

