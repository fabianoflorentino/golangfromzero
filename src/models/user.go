package models

import (
	"net/mail"
	"strings"
	"time"

	"github.com/fabianoflorentino/golangfromzero/src/helper"
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
	if err := u.isNameValid(); err != nil {
		return err
	}

	if err := u.isEmailValid(); err != nil {
		return err
	}

	if u.isNewRegister(register) && u.Password == "" {
		return ErrPasswordBlank
	}

	if err := u.hashPasswd(u.Password); err != nil {
		return err
	}

	return nil
}

func (u *User) isNameValid() error {
	n := strings.TrimSpace(u.Name)

	if n == "" {
		return ErrNameBlank
	}

	return nil
}

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

func (u *User) hashPasswd(register string) error {
	if register == "new" {
		p := strings.TrimSpace(u.Password)
		hashedPasswd, err := helper.DoPasswdHash(p)
		if err != nil {
			return err
		}

		u.Password = string(hashedPasswd)
	}

	return nil
}

// isNewRegister validate if registry is new.
func (u *User) isNewRegister(register string) bool {
	if register == "new" {
		return true
	}

	return false
}
