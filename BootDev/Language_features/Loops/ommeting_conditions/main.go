package main

func maxMessages(thresh int) int {
	msgCounter := 0
	totalCost := 0
	for i:=0; ; i++{ //skipping the condition and moving inside the loop
		currentMsgCost := 100+i
	//The look before you leap method
		if totalCost + currentMsgCost > thresh { //cheaking if cost is < thresh
			break
		}
		totalCost += currentMsgCost
		msgCounter++ //updating the counter at the very last.
	}
	return msgCounter
}
