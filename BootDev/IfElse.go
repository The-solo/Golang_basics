package main 

import (
    "fmt"
)

// In go we don't use brackets for condition but curly brackets are used on the same line.

func main(){
    //Declaring the length inside the if to limit it's scope. Has no use outside.
    if length := 69; length < 33{
        fmt.Println("Length is shorter.")
    } else {
        fmt.Println("The length is greater.")
    }
}
