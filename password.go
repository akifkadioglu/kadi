package kadi

import "golang.org/x/crypto/bcrypt"

// Hash returns hashed password
func Hash(value string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)
	return string(hashedPassword)
}

// CompareHash compares hashed password with value
func CompareHash(hashedString, valueString string) error {
	hashed := []byte(hashedString)
	value := []byte(valueString)
	err := bcrypt.CompareHashAndPassword(hashed, value)
	return err
}
