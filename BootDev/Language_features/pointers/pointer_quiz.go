package main

import ("fmt")

func main() {
	var x int = 50
	var y *int = &x //y is designed to store pointer to a data of type int.
	// which is exactly what we just did.
	*y = 100 // following the pointer and chaning value frim 50 to 100.

	fmt.Println(&y) // address of y. AKA pointer.
	fmt.Println(*y) // dereferning operator.
	fmt.Println(x)
}
