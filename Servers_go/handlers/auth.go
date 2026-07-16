package handlers

import (
	"os"
	"log"
	"time"
//	"errors"
	"net/http"
	"encoding/json"	
    "github.com/google/uuid"
	"server_basics.com/internal/auth"

)

type LoginInfo struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email"`
	Token	  string 	`json:"token`
}

var Jwt_secret = os.Getenv("JWT_SECRET")

func (state *ApiCfgState) Login(w http.ResponseWriter, req *http.Request) {

	type reqParam struct{
		Email  string `json:"email"`
		Password  string `json:password"`
		Expires_In_Seconds int `json:"expires_in_seconds default:"3600"`
	}

	decoder := json.NewDecoder(req.Body)
	param := reqParam{}
	defer req.Body.Close()

	err := decoder.Decode(&param) //stictly pass by reference.
	if err != nil {
		log.Printf("Error decoding the request body: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := state.DB.GetUser(req.Context(), param.Email)
	if err != nil {
		log.Printf("User does not found!!")
		w.WriteHeader(401)
	}
	isRight, _ := auth.CheckPassword(param.Password, user.HashedPassword)// password checking.

	duration := time.Duration(param.Expires_In_Seconds) * time.Second // casting int to time in seconds.

	//generating the JWT token.
	token, err := auth.MakeJWT(user.ID, Jwt_secret, duration)
	if err != nil {
		log.Printf("error generating token")
		w.WriteHeader(500)
	}

	if isRight {
		resBody := LoginInfo{
			ID : user.ID,
			CreatedAt : user.CreatedAt.Time,
			UpdatedAt : user.UpdatedAt.Time,
			Email : user.Email,
			Token : token,
		}	
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(resBody)

	} else {
		w.WriteHeader(401)
		log.Printf("User doesn't exist")
	}
}

