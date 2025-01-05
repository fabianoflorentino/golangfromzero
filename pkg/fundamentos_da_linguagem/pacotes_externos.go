// Package sendmail provides utilities for validating email addresses.
// It uses the checkmail library to verify the format of the email addresses.
package fundamentos_da_linguagem

import (
	"github.com/badoux/checkmail"
)

// CheckEmail validates the format of the provided email address.
// It returns "Email válido!" if the email format is valid, otherwise it returns "Email inválido!".
//
// Parameters:
//
//	email - the email address to be validated
//
// Returns:
//
//	A string indicating whether the email format is valid or not.
func PacoteExterno(email string) string {
	if checkmail.ValidateFormat(email) != nil {
		return "Email inválido!"
	}

	return "Email válido!"
}
