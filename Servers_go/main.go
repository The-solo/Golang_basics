package main

import(
	"io"
	"fmt"
	"net/http"
)


func main(){

	mux := http.NewServeMux() //Http request multiplexer

	mux.Handle("/app/",http.StripPrefix("/app/", http.FileServer(http.Dir(".")))) // serving files from the current dir (index.html)

	// we write the response to ResponseWriter.
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, req *http.Request){ // this method satisfies the http.Handler interface
		w.Header().Set("Content-Type", "text/plain; charset=utf-8") // obvious headers.
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, "The server is up & ready to server.")
	})

	port := "8080"
	server := &http.Server{
		Addr : ":"+port,
		Handler : mux,
	}
	fmt.Println("The server is up & running on port"+server.Addr)
	
 	err := http.ListenAndServe(server.Addr, server.Handler)
	if err != nil {
		 fmt.Printf("Server connection failed", err)
	}
}
