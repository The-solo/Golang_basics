package main

import("fmt")

func main(){

    defer fmt.Println("world")
// The execution is delayed till the below func executes.

    fmt.Println("Hello")

}
