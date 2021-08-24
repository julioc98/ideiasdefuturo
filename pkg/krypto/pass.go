package krypto

import "golang.org/x/crypto/bcrypt"

// Hash implements root.Hash
type Hash struct{}

// Encrypt a salted hash for the input string
func (c *Hash) Encrypt(s string) string {
	saltedBytes := []byte(s)

	hashedBytes, _ := bcrypt.GenerateFromPassword(saltedBytes, bcrypt.DefaultCost)

	hash := string(hashedBytes[:])
	return hash
}

// Compare string to generated hash
func (c *Hash) Compare(hash string, s string) error {
	incoming := []byte(s)
	existing := []byte(hash)
	return bcrypt.CompareHashAndPassword(existing, incoming)
}
