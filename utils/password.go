package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword generates a hashed password from a given password string.
//
// password is the string to be hashed.
// Returns a string representing the hashed password and an error if the hashing process fails.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash checks if a given password matches a hashed password.
//
// password is the string to be checked, and hash is the hashed password string.
// Returns a boolean indicating whether the password matches the hash.
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
