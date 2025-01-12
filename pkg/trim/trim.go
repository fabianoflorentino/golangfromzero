// Package trim provides functionality to remove leading and trailing white space characters from strings.
// It defines a spaceTrimmer type with a method to trim white spaces from a given string.
package trim

import "strings"

// spaceTrimmer is a struct that provides methods to trim spaces from strings.
// It does not contain any fields and serves as a receiver for methods related to space trimming.
type spaceTrimmer struct{}

// New creates and returns a new instance of spaceTrimmer.
// This function initializes a spaceTrimmer struct and returns a pointer to it.
func New() *spaceTrimmer {
	return &spaceTrimmer{}
}

// String trims all leading and trailing white space characters from the input string.
// It takes a single string argument 'txt' and returns a new string with the white spaces removed.
func (s spaceTrimmer) String(txt string) string {
	return strings.TrimSpace(txt)
}
