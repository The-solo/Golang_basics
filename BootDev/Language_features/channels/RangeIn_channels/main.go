package main

func concurrentFib(n int) []int {
	integers := make(chan int)
	value := []int{}
	go fibonacci(n, integers) //making the fibonacci func concurrent.
	for item := range integers{
		value = append(value, item)
		//exits only when the channel is closed.
	}
	return value
}

// The ranging over channel blocks for every iteration to check if item is there.

// don't touch below this line

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch) //closing the channel is crucial.
}
