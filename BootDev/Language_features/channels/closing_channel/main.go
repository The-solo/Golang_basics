package main

func countReports(numSentCh chan int) int {
	total := 0
	for {
		ch, ok := <-numSentCh // ch contains the actul data of type int that we pass
		if !ok { //introducing the breaking condition.
			break
		}
		total+=ch
	}
	return total
}

// don't touch below this line

func sendReports(numBatches int, ch chan int) {
	for i := 0; i < numBatches; i++ {
		numReports := i*23 + 32%17
		ch <- numReports
	}
	close(ch)
}
