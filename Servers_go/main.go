package main

import(
	"fmt"
	"net/http"
)


func main(){

	mux := http.NewServeMux() //Http request multiplexer

	mux.Handle("/",http.FileServer(http.Dir("."))) // serving files from the current dir (index.html)

	//standard http server struct.
	port := "8080"
	server := &http.Server{
		Addr : ":"+port,
		Handler : mux,
	}
	fmt.Println("The server is up & running on localhost:8080...")

 	err := http.ListenAndServe(server.Addr, server.Handler)
	if err != nil {
		 fmt.Printf("Server connection failed", err)
	}
}
