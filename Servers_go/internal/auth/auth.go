package auth

import (
	"log"
	"time"
	"github.com/google/uuid"
	"github.com/alexedwards/argon2id"
	"github.com/golang-jwt/jwt/v5"
)

func HashPassword (password string) (string, error) {
	hash, err := argon2id.CreateHash(password, argon2id.DefaultParams)
	if err != nil {
		log.Fatal(err)
		return password, err
	}
	return hash, nil
}

func CheckPassword (password, hash string) (bool, error) {
	_, err := argon2id.ComparePasswordAndHash(password, hash)
	if err != nil {
		log.Fatal(err)
		return false, err
	}
	return true, nil
}



 type Claims struct {
	 jwt.RegisteredClaims
 }

func MakeJWT(userID uuid.UUID, tokenSecret string, expiresIn time.Duration) (string, error) {

	// Basics of validation using claims.
	claimsData := &Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
			ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(expiresIn)),
			Issuer:    "chirpy-access",
			Subject:   userID.String(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsData) // Generate token.
	tokenString, err := token.SignedString(tokenSecret)

	if err != nil{
		log.Print("Error signing key", err)
		return "", err
	}
	return tokenString, nil
}



//JWT token validation.
func ValidateJWT(tokenString, tokenSecret string) (uuid.UUID, error) {	

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (any, error) {
		return []byte(tokenSecret), nil
	})

	if err != nil {
		log.Fatal(err)
		return uuid.Nil, err // nil uuid return type

	}
	if claims, ok := token.Claims.(*Claims); ok {

		subjectStr, err := claims.GetSubject() // Contains the userId/uuid
		if err != nil {
			return uuid.Nil, err
		}
		return uuid.Parse(subjectStr) // the GetSubject method return 2 values including error.
	} 

	return uuid.Nil, err
}


