package helper

import "golang.org/x/crypto/bcrypt"

// DoPWDDoPasswdHashHash returns the bcrypt hash of the password at the given cost
func DoPasswdHash(pwd string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
}

// IsPasswdValid compares a bcrypt hashed password with its possible plaintext equivalent.
// Returns nil on success, or an error on failure.
func IsPasswdValid(pwdHashed, pwd string) error {
	return bcrypt.CompareHashAndPassword([]byte(pwdHashed), []byte(pwd))
}
