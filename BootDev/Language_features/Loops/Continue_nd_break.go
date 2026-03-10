package main

import("fmt")



func printPrimes(max int) {
	for i:=2; i<=max; i++{
		isPrime := true
//cheaking if i is divisible by any number between 2 to i-1
		for j:=2; j<i; j++ {
			if i % j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Println(i)
		}
	}
}

// don't edit below this line
func test(max int) {
	fmt.Printf("Primes up to %v:\n", max)
	printPrimes(max)
	fmt.Println("===============================================================")
}

func main() {
	test(10)
	test(20)
	test(30)
}



	
