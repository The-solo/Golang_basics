package auth

import (
	"log"
	"github.com/alexedwards/argon2id"
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
		return false, err
	}
	return true, nil
}




