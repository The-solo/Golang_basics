package handlers

import(
	"log"
	"time"
	//"context"
	"net/http"
	"encoding/json"	
	"github.com/google/uuid"
	"server_basics.com/internal/database"

)
// We imported database.Queries directly. 
// we made this struct to get access to db and attach below methods to it.
type ApiCfgState struct {
	DB *database.Queries
}

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email"`
}

// func decodeBody(reqParam{}, w http.ResponseWriter, req *http.Request) (param, error){ 
//
// 	decoder := json.NewDecoder(req.Body)
// 	param := reqParam{}
// 	defer req.Body.Close()
//
// 	err := decoder.Decode(&reqParam)
// 	if err != nil {
// 		log.Printf("Error decoding the request body: %s", err)
// 		w.WriteHeader(http.StatusBadRequest) // 400 is safer for bad JSON input
// 		return nil, err
// 	}
// 	return param, nil 
// }

func (state *ApiCfgState)CreateUserHandler(w http.ResponseWriter, req *http.Request) {

	type reqParam struct {
		Email string `json:"email"`
	}

	decoder := json.NewDecoder(req.Body)
	param := reqParam{}
	defer req.Body.Close()

	err := decoder.Decode(&param)
	if err != nil {
		log.Printf("Error decoding the request body: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// 4. Use the local DB copy to execute your sqlc generated query.
	user, err := state.DB.CreateUser(req.Context(), param.Email)//using context directly.
	if err != nil {
		log.Printf("Error creating user in DB: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resBody := User{
		ID : user.ID,
		CreatedAt : user.CreatedAt.Time,
		UpdatedAt : user.UpdatedAt.Time,
		Email : user.Email,
	}

	data, err := json.Marshal(resBody)
	if err != nil{
		log.Printf("Error marshaling the json %s", err)
		w.WriteHeader(500)
		return
	}
	// 5. Send back the automatically generated sqlc user object
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
}


