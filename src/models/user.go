package models

import (
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
func (u *User) Validate() error {
	if err := u.isBlank(); err != nil {
		return err
	}

	u.trimSpace()

	return nil
}

func (u *User) isBlank() error {
	if u.Name == "" {
		return ErrNameBlank
	}

	if u.Email == "" {
		return ErrEmailBlank
	}

	if u.Password == "" {
		return ErrPasswordBlank
	}

	return nil
}

func (u *User) trimSpace() {
	strings.TrimSpace(u.Name)
	strings.TrimSpace(u.Email)
}
