package models

import (
	"net/mail"
	"strings"
	"time"

	"github.com/google/uuid"
)

// User represents a user in the system.
type User struct {
	ID        uuid.UUID `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

// Validate check if any field is blank, and with all leading and trailing white space removed
func (u *User) Validate(register string) error {
	u.trimSpace()

	if err := u.isBlank(); err != nil {
		return err
	}

	if err := u.isValidEmail(u.Email); err != nil {
		return err
	}

	if u.isNewRegister(register) && u.Password == "" {
		return ErrPasswordBlank
	}

	return nil
}

func (u *User) isBlank() error {
	if u.Name == "" {
		return ErrNameBlank
	}

	if u.Email == "" {
		return ErrEmailBlank
	}

	return nil
}

// trimSpace remove blanks spaces before and after the string.
func (u *User) trimSpace() {
	strings.TrimSpace(u.Name)
	strings.TrimSpace(u.Email)
}

// isNewRegister validate if registry is new.
func (u *User) isNewRegister(register string) bool {
	if register == "new" {
		return true
	}

	return false
}

// isValidEmail parses a single RFC 5322 address using mail.ParseAddress
func (u *User) isValidEmail(email string) error {
	if _, err := mail.ParseAddress(email); err != nil {
		return ErrInvalidEmailFormat
	}

	return nil
}
