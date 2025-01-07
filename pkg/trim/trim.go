// Package trim provides utility functions for trimming strings.
package trim

import "strings"

// TrimString removes all leading and trailing white space characters from the input string.
// It takes a single string argument 'text' and returns a new string with the white space trimmed.
func String(text string) string {
	return strings.TrimSpace(text)
}
