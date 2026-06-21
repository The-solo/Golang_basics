package handlers

import(
	"log"
	//"time"
	//"context"
	"net/http"
	"encoding/json"	
	"server_basics.com/internal/database"

)

// We imported database.Queries directly.
type ApiCfgState struct {
	DB *database.Queries
}

func (state *ApiCfgState)CreateUserHandler(w http.ResponseWriter, req *http.Request) {

	ctx := req.Context()// initializing the context
	type reqParam struct {
		Email string `json:"email"`
	}

	decoder := json.NewDecoder(req.Body)
	params := reqParam{}

	err := decoder.Decode(&params)
	if err != nil {
		log.Printf("Error decoding the request body: %s", err)
		w.WriteHeader(http.StatusBadRequest) // 400 is safer for bad JSON input
		return
	}

	// 4. Use the local DB copy to execute your sqlc generated query.
	user, err := state.DB.CreateUser(ctx, params.Email)
	if err != nil {
		log.Printf("Error creating user in DB: %s", err)
		w.WriteHeader(http.StatusInternalServerError) // 500 for database failures
		return
	}

	// 5. Send back the automatically generated sqlc user object
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

