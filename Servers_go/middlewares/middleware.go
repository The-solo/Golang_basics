package middleware

import(
	"net/http"
	"log"
	"server_basics.com/config"
)

// since we are exporting all thse functions they should be public. as in capital first letter.

func MiddlewareMetricsInc (cfg * config.ApiConfig, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request){ // http.HandlerFunc is used not HandleFunc
		hits := cfg.FileserverHits.Add(1) //Increamenting the request count
		log.Printf("Number of request hits : %v", hits)
		next.ServeHTTP(w, req)
	})
}

func MiddlewareLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
