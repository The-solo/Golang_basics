package handlers

import(
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"server_basics.com/config"
)


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
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		cfg.FileserverHits.Store(0) //resetting the count.
	}
}

func respondWithError(w http.ResponseWriter, code int, msg string){
	w.Header().Set("Content-Type", "text/utf-8")
	w.WriteHeader(code)
	w.Write([]byte(msg)) //write the message
	return
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(payload.([]byte))
	return
}

func ChirpValidator(w http.ResponseWriter, req *http.Request) {	
	type reqParam struct{
		Body string
	}

	decoder := json.NewDecoder(req.Body)
	params := reqParam{}

	err := decoder.Decode(&params)
	if err != nil {
		log.Printf("Error encoding the request body %s", err)
		w.WriteHeader(500)
		return
	} 

	//Making sure the length Chirp is < 140
	if len(params.Body) > 140 {
		respondWithError(w, 400, "Chirp is too long")
		return
	}

	defer req.Body.Close()

	type resValue struct {
		Valid bool
	}

	resBody := resValue{
		Valid : true,
	}

	data, err := json.Marshal(resBody)
	if err != nil{
		log.Printf("Error marshaling the json %s", err)
		w.WriteHeader(500)
		return
	}
	respondWithJSON(w, 200, data)
}
