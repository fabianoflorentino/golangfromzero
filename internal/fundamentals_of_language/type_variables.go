// Package fundamentos_da_linguagem provides fundamental examples and concepts of the Go programming language.
// It includes demonstrations of variable declarations, constants, and formatted string outputs.
package fundamentals_of_language

import "fmt"

// TypeVariables demonstrates the declaration and initialization of different types of variables
// and constants in Go. It returns a formatted string that includes the values and types of these
// variables and constants.
func TypeVariables() string {
	var var1 string = "variable 1"
	var2 := "variable 2"

	var (
		var3 string = "variable 3"
		var4 string = "variable 4"
	)

	var5, var6 := "variable 5", "variable 6"

	const constant1 string = "Constante 1"

	variable := fmt.Sprintf(
		"\nvar1: %s :: %T\nvar2: %s :: %T\nvar3: %s :: %T\nvar4: %s :: %T\nvar5: %s :: %T\nvar6: %s :: %T\nconstant1: %s :: %T",
		var1, var1, var2, var2, var3, var3, var4, var4, var5, var5, var6, var6, constant1, constant1,
	)

	return variable
}
