package main

import (
	"time"
)

// Here the saveBackups is the single gorutine which is listening to multiple channels.
func saveBackups(snapshotTicker, saveAfter <-chan time.Time, logChan chan string) {
	for{
		select{
		case _, ok := <-snapshotTicker:
			if ok {
				takeSnapshot(logChan)
			} else {
				return
			}
		case _, ok := <-saveAfter:
			if ok {
				saveSnapshot(logChan)
				return  // this is actually terminating the loop infinite for loop.
			} else {
				return
			}
		default:
			waitForData(logChan)
			time.Sleep(500 * time.Millisecond)
		}
	}

}

// don't touch below this line

func takeSnapshot(logChan chan string) {
	logChan <- "Taking a backup snapshot..."
}

func saveSnapshot(logChan chan string) {
	logChan <- "All backups saved!"
	close(logChan)
}

func waitForData(logChan chan string) {
	logChan <- "Nothing to do, waiting..."
}

/*
Every channel is bidirection by nature but we can make them read and write only.

// This is the readonly channel function

	func readCh(ch <-chan int) {
    	// ch can only be read from (because of the postion of the arrow)
    	// in this function
	}


// This is the wrie only channel function.

	func writeCh(ch chan<- int) {
    // ch can only be written to (because of the direction  of the Arrow)
    // in this function
}
*/
