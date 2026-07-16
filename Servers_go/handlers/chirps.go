package handlers

import(
	"log"
	"time"
	"context"
	"net/http"
	"encoding/json"	
	"github.com/google/uuid"
	"server_basics.com/internal/auth"
	"server_basics.com/internal/database"
)

func (state *ApiCfgState) CreateChirpHandler (w http.ResponseWriter, req *http.Request){
	ctx,cancle := context.WithTimeout(context.Background(), 10*time.Millisecond)
	/* Making a context with timeout so the connection doesn't hang.
	context.Background() mainly used as a top/ parent level context.
	WithTimeout return deadline context which release the resouces.
	The context essetially controls lifecycle of API, netowork requests and Background operations.*/
	defer cancle()

	type reqParam struct {
		Body string    `json:"body"`
		ID   uuid.UUID `json:"user_id"`
		Token string   `json:"Authorization`
	}

	decoder := json.NewDecoder(req.Body)
	param := reqParam{}
	defer req.Body.Close()

	err := decoder.Decode(&param)
	if err != nil {
		log.Printf("Error decoding the request body: %s", err)
		w.WriteHeader(400)
		return
	}
	//checking if bearer is present ?
	_, err = auth.GetBearerToken(req.Header) // GetBearerToken expects input of type http.Header
	if err != nil {
		log.Printf("Missing bearer!!", err)
		w.WriteHeader(401)
		return
	}
	_, err = auth.ValidateJWT(param.Token, Jwt_secret) //checking if the bearer is valid?
	if err != nil{
		log.Fatal("Invalid jwt token", err)
		w.WriteHeader(401)
		return
	}

	if len(param.Body) > 140 {
		respondWithError(w, 400, "Chirp is too long")
		return
	}

	cleanBody := checkFoulWords(param.Body)
	chirp, err := state.DB.CreateChirp(ctx,
		database.CreateChirpParams{ 		
		Body:   cleanBody,
		UserID: param.ID,
	},)

	if err != nil {
		log.Printf("Error creating chirp in DB please try again : %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(chirp)
}

//Get all chirps.
func (state *ApiCfgState) ReturnAllChirp (w http.ResponseWriter, req *http.Request){

	chirp, err := state.DB.GetAllChirps(req.Context())
	if err != nil{
		log.Printf("error returning chirps %s", err)
		w.WriteHeader(500)
		return 
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(chirp)
}

//get single chirp by {id}
func (state *ApiCfgState) GetChirp (w http.ResponseWriter, req *http.Request){

	StringId := req.PathValue("id") //Reading the PathValue 'id'

	chirpID, err := uuid.Parse(StringId) //converting string input into uuid.
	if err != nil {
		log.Printf("Invalid chirp ID format %s", err)
		w.WriteHeader(500)
		return
	}
	
	chirp, err := state.DB.GetChirp(req.Context(), chirpID)
	if err != nil{
		log.Printf("error returning chirp %s", err)
		w.WriteHeader(500)
		return 
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(chirp)

}

