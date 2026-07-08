package handlers

import (
	"log"
//	"errors"
	"net/http"
	"encoding/json"	
	"github.com/google/uuid"
	"github.com/golang-jwt/jwt/v5"
	"server_basics.com/internal/auth"

)

func (state *ApiCfgState) Login(w http.ResponseWriter, req *http.Request) {

	type reqParam struct{
		Email  string `json:"email"`
		Password  string `json:password`
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

	if isRight {
		resBody := User{
			ID : user.ID,
			CreatedAt : user.CreatedAt.Time,
			UpdatedAt : user.UpdatedAt.Time,
			Email : user.Email,
		}	
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(resBody)

	} else {
		w.WriteHeader(401)
		log.Printf("User doesn't exist")
	}
}


//JWT token validation.
func ValidateJWT(tokenString, tokenSecret string) (uuid.UUID, error) {	

	token, err := jwt.ParseWithClaims(tokenString, &auth.Claims{}, func(token *jwt.Token) (any, error) {
		return []byte(tokenSecret), nil
	})

	if err != nil {
		log.Fatal(err)
		return uuid.Nil, err // nil uuid return type

	}
	if claims, ok := token.Claims.(*auth.Claims); ok {

		subjectStr, err := claims.GetSubject() // Contains the userId/uuid
		if err != nil {
			return uuid.Nil, err
		}
		return uuid.Parse(subjectStr) // the GetSubject method return 2 values inclduing error.
	} 

	return uuid.Nil, err
}



