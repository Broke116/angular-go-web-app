package security

import (
	"strings"

	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

// HashInterface is an interface for password hashing
type HashInterface interface {
	Generate(s string) (string, error)
	Compare(hash string, s string) error
}

// Hash implements Hash interfaces
type Hash struct{}

var separator = "||"

// Generate is used to generate a cypted password
func (h *Hash) Generate(s string) string {
	salt := uuid.NewV1().String()
	saltedBytes := []byte(s + salt)
	hashedBytes, err := bcrypt.GenerateFromPassword(saltedBytes, bcrypt.DefaultCost)

	if err != nil {
		return ""
	}

	hash := string(hashedBytes[:])
	return hash + separator + salt
}

// Compare is used to check the hashed password equilavent of the entered plain text
func (h *Hash) Compare(hash string, s string) error {
	parts := strings.Split(hash, separator)

	if len(parts) != 2 {
		// TO DO improve the error logic
	}

	input := []byte(s + parts[1])
	existing := []byte(parts[0])
	return bcrypt.CompareHashAndPassword(existing, input)
}
