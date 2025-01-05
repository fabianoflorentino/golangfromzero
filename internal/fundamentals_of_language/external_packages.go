// Package sendmail provides utilities for validating email addresses.
// It uses the checkmail library to verify the format of the email addresses.
package fundamentals_of_language

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
func ExternalPackages(email string) string {
	if checkmail.ValidateFormat(email) != nil {
		return "mail invalid!"
	}

	return "mail valid!"
}
