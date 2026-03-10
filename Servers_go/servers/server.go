package main

import(
	"fmt"
	"net/http"
)


type Handler interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}

func main(){

	mux := http.NewServeMux() //Http request multiplexer

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	type Server struct{
		Addr string 
		handler Handler
	}
	server := Server{Addr: ":8080", handler: mux}

	fmt.Println("The server is up & running on localhost:8080...")
	http.ListenAndServe(server.Addr, server.handler)
}
