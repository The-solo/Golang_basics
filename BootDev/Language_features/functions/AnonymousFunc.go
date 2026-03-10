package main 

import ("fmt")

func main(){

    char := "me"
    x := len(char)
    fmt.Println(x)

    double := func (a int) int {
        return a*2
    }

    fmt.Println(double(10))


// Anonymous func
    func(a int) int {
        return a+a
    } (x) //called right after.
}
