package handlers

import(
	"io"
	"fmt"
	"log"
	"strings"
	"net/http"
	"server_basics.com/config"
)

func ServerHealthCheck(w http.ResponseWriter, req * http.Request){
	w.Header().Set("Content-Type", "text/plain; charset=utf-8") // obvious headers.
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "The server is up & ready to server.")
}

// These are now regular functions with closure because the struct is in another file.
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

		ctx := req.Context()// initializing the context
		if cfg.Platform != "dev"{ // Dangerours method allowed in local development only.
			w.WriteHeader(403)
			return 
		}

		err := cfg.Database.DeleteUser(ctx) //deleting the database.
		if err != nil{
			log.Printf("error deleting the database: %s",err)
		}
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		cfg.FileserverHits.Store(0) //resetting the count.
	}
}

//function for the error response.
func respondWithError(w http.ResponseWriter, code int, msg string){
	w.Header().Set("Content-Type", "text/utf-8")
	w.WriteHeader(code)
	w.Write([]byte(msg)) //write the message
	return
}

//function for JsonResponse.
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) { // Changed interface{} to []byte
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write((payload).([]byte))
	return
}

func checkFoulWords(body string) string{
	restricted := [3]string{"kerfuffle", "sharbert", "fornax"}
	newSlice := strings.Split(body, " ") //returns slice of words seperated by space.

	for i := range newSlice{
		for _, foulWord := range restricted{
			if strings.ToLower(newSlice[i]) == foulWord{
				newSlice[i] = "****"
			}
		}
	}
	return strings.Join(newSlice, " ") // rejoining the cleaned slice.
}

