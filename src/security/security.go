package security

import "golang.org/x/crypto/bcrypt"

// Hash generates a bcrypt hash of the password.
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// VerifyPassword compares a bcrypt hashed password with its possible plaintext equivalent.
func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
