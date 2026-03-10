package main

import (
	"sync"
	"time"
)

type safeCounter struct {
	counts map[string]int
	mu     *sync.Mutex // To make sure operation on the same instance of the mutex.
}

func (sc safeCounter) inc(key string) {
	func(){ // It's a good practice to structure the protected code inside the function(NOT NEEDED HERE)
		sc.mu.Lock() //lock the mutex for the rest of the function.
		defer sc.mu.Unlock()//unlock it by the end.
		sc.slowIncrement(key)
	}()
}

func (sc safeCounter) val(key string) int {
	sc.mu.Lock() 
	defer sc.mu.Unlock()	
	return sc.slowVal(key)
}

// don't touch below this line

func (sc safeCounter) slowIncrement(key string) {
	tempCounter := sc.counts[key]
	time.Sleep(time.Microsecond)
	tempCounter++
	sc.counts[key] = tempCounter
}

func (sc safeCounter) slowVal(key string) int {
	time.Sleep(time.Microsecond)
	return sc.counts[key]
}
