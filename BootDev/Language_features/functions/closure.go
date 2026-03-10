package main 

import ("fmt")

func main(){

    fmt.Println(adder()(5))

    caller := adder()
   // baller := adder() 
// Here the each instance of the adder is it's own seperate entity.
// Eg. The baller and caller both are seperate.
    fmt.Println(caller(5))

}

func adder() func(int) int {
	var sum int
	return func(input int) int{ //closure starts here.
		sum += input
		return sum
	}
}
// The sum variable is persistant as long as adder() exists
