package main

import(
	"io"
	"fmt"
	"net/http"
	"server_basics.com/middlewares" //importing the middlewares
)

func main(){

	router := http.NewServeMux() //Http request multiplexer/router

	apiCfg := &middleware.ApiConfig{}// making the local copy.
	
	// serving files from the current dir (index.html)
	router.Handle("/app/",http.StripPrefix("/app/", http.FileServer(http.Dir("."))))
	
	// we write the response to ResponseWriter.
	router.HandleFunc("/healthz", func(w http.ResponseWriter, req *http.Request){
		w.Header().Set("Content-Type", "text/plain; charset=utf-8") // obvious headers.
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, "The server is up & ready to server.")
	})

	router.HandleFunc("/metrics",(apiCfg.Metric)) //metric router.
	router.HandleFunc("/reset",(apiCfg.Reset)) //metric reset router
	
	port := "8080"
	server := &http.Server{
		Addr : ":"+port,
		Handler : router,
	}
	fmt.Println("The server is up & running on port"+server.Addr+"....")
	
 	err := http.ListenAndServe(server.Addr, 
		middleware.MiddlewareLog(
			apiCfg.MiddlewareMetricsInc(server.Handler), //wrapped the handler inside the count middlware.
	)) // logging every request handler using middlware.

	if err != nil {
		 fmt.Printf("Server connection failed %v", err)
	}
}
