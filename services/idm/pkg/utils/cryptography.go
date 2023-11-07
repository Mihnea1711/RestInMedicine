package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes a user's password using bcrypt
func HashPassword(password string) (string, error) {
	// Generate a salt for the password
	salt, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("[IDM] Failed to hash the password: %v", err)
		return "", err
	}

	log.Println("[IDM] Password hashed successfully.")
	return string(salt), nil
}

// VerifyPassword compares the user's input with the hashed password
func VerifyPassword(hashedPassword, userInputPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(userInputPassword))
	if err != nil {
		log.Printf("[IDM] Password verification failed: %v", err)
		return err
	}

	log.Println("[IDM] Password verification succeeded.")
	return nil
}
