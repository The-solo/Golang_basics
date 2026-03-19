package handlers

import(
	"fmt"
	"net/http"
	"server_basics.com/config"
)


func Metric(cfg * config.ApiConfig) http.HandlerFunc{
	return func(w http.ResponseWriter, req *http.Request){
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		//	log.Printf("Hits:%s", cfg.fileserverHits.Load())//printing the requests count
		htmlContent := fmt.Sprintf(`<html>
		<body>
		<h1>Welcome to the Admin pannel</h1>
		<p>The server has been visited %d times!</p>
		</body>
		</html>`, cfg.FileserverHits.Load())

		fmt.Fprint(w, htmlContent)  // w/io.writer -> destination and content returns the no. of bytes.
	}
} 
func Reset (cfg * config.ApiConfig) http.HandlerFunc{
	return func(w http.ResponseWriter, req *http.Request){
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		cfg.FileserverHits.Store(0) //resetting the count.
	}
}
