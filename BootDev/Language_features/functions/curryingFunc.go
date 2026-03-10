package main

import ("fmt")

func main(){

    fmt.Println(add(2)(3))
    fmt.Println(whatever(100, divide)) // prints the divide func itself.
    fmt.Println(whatever(100, divide)(5)) //prints result = 20
    return

}

//Currying func examples.


func add(a int) func(int)int{
    return func(b int) int{
        return a*b
    }
}


func divide(a, b int)int{
    a += 10 
    b += 5
    return a/b
}


//               type signature of divide func. The return type.
func whatever(a int, divide func(int, int)int) func(int) int{
    return func(b int) int{
        return divide(a, b)
    }
}
