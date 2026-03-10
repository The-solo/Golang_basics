package main

import (
	"time"
)

type email struct {
	body string
	date time.Time
}

func checkEmailAge(emails [3]email) [3]bool {
	isOldChan := make(chan bool)

	go sendIsOld(isOldChan, emails) //spawn the gorutine to make it concurrent. 
	//Even though the data is moving the communication is blocking/synchronus
	// The next one cannot be processed until the first one finishes.
	isOld := [3]bool{}
	isOld[0] = <-isOldChan //get's whichever one finishes first.
	isOld[1] = <-isOldChan //get's whichever one finishes second.
	isOld[2] = <-isOldChan //get's whichever one finishes third.
	return isOld
}

// don't touch below this line

func sendIsOld(isOldChan chan<- bool, emails [3]email) {
	for _, e := range emails {
		if e.date.Before(time.Date(2020, 0, 0, 0, 0, 0, 0, time.UTC)) {
			isOldChan <- true
			continue
		}
		isOldChan <- false
	}
}
