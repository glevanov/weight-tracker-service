package auth

import (
	"crypto/rand"
	"encoding/hex"

	"golang.org/x/crypto/scrypt"
)

// HashPassword hashes a password with a given salt using scrypt
// Matches Node.js scrypt defaults: N=16384, r=8, p=1, keyLen=64
func HashPassword(password, salt string) (string, error) {
	key, err := scrypt.Key([]byte(password), []byte(salt), 16384, 8, 1, 64)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(key), nil
}

// GenerateSalt generates a random 16-byte salt and returns it as a hex string
func GenerateSalt() (string, error) {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(salt), nil
}
