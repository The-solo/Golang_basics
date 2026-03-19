package middleware

import(
	"fmt"
	"net/http"
	"log"
	"sync/atomic"
)

// since we are exporting all thse functions they should be public. as in capital first letter.

type ApiConfig struct {
	fileserverHits atomic.Int32
}

func (cfg *ApiConfig) MiddlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request){ // http.HandlerFunc is used not HandleFunc
		hits := cfg.fileserverHits.Add(1) //Increamenting the request count
		log.Printf("Number of request hits : %v", hits)
		next.ServeHTTP(w, req)
	})
}

func (cfg * ApiConfig) Metric(w http.ResponseWriter, req *http.Request){
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	//	log.Printf("Hits:%s", cfg.fileserverHits.Load())//printing the requests count
	htmlContent := fmt.Sprintf(`<html>
	<body>
	<h1>Welcome to the Admin pannel</h1>
	<p>The server has been visited %d times!</p>
	</body>
	</html>`, cfg.fileserverHits.Load())

	fmt.Fprint(w, htmlContent)  // w/io.writer -> destination and content returns the no. of bytes.
}

func (cfg * ApiConfig) Reset(w http.ResponseWriter, req *http.Request){
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	cfg.fileserverHits.Store(0) //resetting the count.
}


func MiddlewareLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
