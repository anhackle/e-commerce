package crypto

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
)

func GetHash(key string) string {
	hash := sha256.New()
	hash.Write([]byte(key))
	hashBytes := hash.Sum(nil)

	return hex.EncodeToString(hashBytes)
}

func GenerateSalt(length int) (string, error) {
	salt := make([]byte, length)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	return hex.EncodeToString(salt), nil
}

func HashPassword(password string, salt string) string {
	saltedPassword := password + salt
	hashedPassword := sha256.Sum256(([]byte(saltedPassword)))
	return hex.EncodeToString(hashedPassword[:])
}

func MatchingPassword(storedHashPassword, password, salt string) bool {
	hashPassword := HashPassword(password, salt)
	return storedHashPassword == hashPassword
}
