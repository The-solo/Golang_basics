package handlers

import(
	"io"
	"fmt"
	"log"
	"strings"
	"net/http"
	"encoding/json"
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
	newSlice := strings.Split(body, " ") //slice of words seperated by space.

	for i := range newSlice{
		for _, foulWord := range restricted{
			if strings.ToLower(newSlice[i]) == foulWord{
				newSlice[i] = "****"
			}
		}
	}
	return strings.Join(newSlice, " ") // rejoining the cleaned slice.
}

func ChirpValidator(w http.ResponseWriter, req *http.Request) {	
	type reqParam struct{
		Body string `json:"chirp`
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
		Output string 
	}

	resBody := resValue{ 
		Valid : true, // making fields public coz they need to be exported to be used by "encoding/json" 
		Output : checkFoulWords(params.Body), // which is an external code.
	}

	data, err := json.Marshal(resBody)
	if err != nil{
		log.Printf("Error marshaling the json %s", err)
		w.WriteHeader(500)
		return
	}
	respondWithJSON(w, 200, data)	
}
