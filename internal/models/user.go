package models

import (
	"net/mail"
	"strings"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// User represents a user in the system.
type User struct {
	ID        uuid.UUID `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

type ValidationMode int

const (
	ValidationCreate ValidationMode = iota
	ValidationUpdate
)

// Validate check if any field is blank, and with all leading and trailing white space removed
func (u *User) Validate(mode ValidationMode) error {
	if err := u.isNameValid(); err != nil {
		return err
	}

	if err := u.isEmailValid(); err != nil {
		return err
	}

	if u.isNewRegister(mode) && u.Password == "" {
		return ErrPasswordBlank
	}

	if err := u.hashPasswd("new"); err != nil {
		return err
	}

	return nil
}

// isNameValid trim spaces from string input and validate if the user name isn't blank
func (u *User) isNameValid() error {
	n := strings.TrimSpace(u.Name)

	if n == "" {
		return ErrNameBlank
	}

	return nil
}

// isEmailValid trim spaces from string input and validate if the user email isn't blank
// and parses a single RFC 5322 address
func (u *User) isEmailValid() error {
	e := strings.TrimSpace(u.Email)

	if e == "" {
		return ErrEmailBlank
	}

	if _, err := mail.ParseAddress(u.Email); err != nil {
		return ErrInvalidEmailFormat
	}

	return nil
}

// hashPasswd trim spaces from string input and create a hash from it when registry is new
func (u *User) hashPasswd(register string) error {
	if register == "new" {
		p := strings.TrimSpace(u.Password)

		hashedPasswd, err := u.doPasswdHash(p)
		if err != nil {
			return err
		}

		u.Password = string(hashedPasswd)
	}

	return nil
}

// isNewRegister validate if registry is new.
func (u *User) isNewRegister(mode ValidationMode) bool {
	return mode == ValidationCreate
}

// doPasswdHash returns the bcrypt hash of the password at the given cost
func (u *User) doPasswdHash(pwd string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
}

// isPasswdValid compares a bcrypt hashed password with its possible plaintext equivalent.
// Returns nil on success, or an error on failure.
func (u *User) isPasswdValid(pwdHashed, pwd string) error {
	return bcrypt.CompareHashAndPassword([]byte(pwdHashed), []byte(pwd))
}
