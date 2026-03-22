package main

import(
	"fmt"
	"net/http"
	"server_basics.com/middlewares"
	"server_basics.com/handlers"
	"server_basics.com/config"
)

func main(){

	router := http.NewServeMux() //Http request multiplexer/router

	apiCfg := &config.ApiConfig{}// making the local copy.
	
	// serving files from the current dir (index.html)
	router.Handle("/app/",http.StripPrefix("/app/", http.FileServer(http.Dir("."))))
	
	// we write the response to ResponseWriter.
	router.HandleFunc("GET /api/healthz",handlers.ServerHealthCheck) 
	router.HandleFunc("GET /admin/metrics",handlers.Metric(apiCfg)) //normal func & closure mechanism. 
	router.HandleFunc("POST /admin/reset",handlers.Reset(apiCfg))
	router.HandleFunc("POST /api/validate_chirp",handlers.ChirpValidator)

	
	port := "8080"
	server := &http.Server{
		Addr : ":"+port,
		Handler : router,
	}
	fmt.Println("The server is up & running on port"+server.Addr+"....")
	
 	err := http.ListenAndServe(server.Addr, 
		middleware.MiddlewareLog(
			middleware.MiddlewareMetricsInc(apiCfg, server.Handler), //wrapped the handler inside the count middlware.
	)) // logging every request handler using middlware.

	if err != nil {
		 fmt.Printf("Server connection failed %v", err)
	}
}
