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


// Generate the JWT token
func MakeJWT(userID uuid.UUID, tokenSecret string, expiresIn time.Duration) (string, error) {

	// Basics of validation using claims.
	claims := &jwt.RegisteredClaims{
		IssuedAt :  jwt.NewNumericDate(time.Now()),
		ExpiresAt : jwt.NewNumericDate(time.Now().Add(expiresIn)),
		Issuer :    "chirpy-access",
		Subject : userID.String(), // converting UUID to string.
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // Generate token.
	jwt, err := token.SignedString(tokenSecret)
	if err != nil{
		log.Print("Error signing key", err)
		return "", err
	}
	return jwt, nil
}
