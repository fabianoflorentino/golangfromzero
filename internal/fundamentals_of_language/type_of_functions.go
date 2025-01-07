// Package fundamentals_of_language demonstrates different types of functions in Go,
// including functions with a single return value, functions with multiple return values,
// and functions assigned to variables. It also includes examples of how to call these
// functions and print their results.
package fundamentals_of_language

import (
	"fmt"

	"github.com/fabianoflorentino/golangfromzero/pkg/trim"
)

// sum takes two integers as input and returns their sum.
// It adds the two input integers and returns the result.
var (
	funcExampleWithReturn = `
// sum takes two integers, num1 and num2, and returns their sum.
// It adds the two input integers and returns the result.
// The function has a single return value.
func sum(num1, num2 int) int {
	return num1 + num2
}
`

	funcExampleWithMoreThanOneReturn = `
// sumAndsub takes two integers as input and returns two integers.
// The first returned integer is the sum of the input integers,
// and the second returned integer is the difference (num1 - num2) of the input integers.
func sumAndsub(num1, num2 int) (int, int) {
	return num1 + num2, num1 - num2
}
`

	funcExampleWithVariable = `
// functionVariable is an anonymous function assigned to a variable.
// It takes a string parameter 'text' and returns the same string.
// This can be used to demonstrate how functions can be assigned to variables in Go.
var functionVariable = func(text string) string {
	return text
}
`
)

// TypeOfFunctions demonstrates different types of functions in Go.
// It includes examples of functions with a single return value, multiple return values,
// and functions assigned to variables. The function prints the results of these examples
// using the fmt.Printf function.
func TypeOfFunctions() {

	resSum := sum(1, 2)
	fmt.Printf("\n%v\nexample of result: %v\n", trim.String(funcExampleWithReturn), resSum)

	resultSum, resultSub := sumAndsub(10, 10)
	fmt.Printf("\n%v\nexample of result: %v, %v", trim.String(funcExampleWithMoreThanOneReturn), resultSum, resultSub)

	funcVariableResult := functionVariable("Hey, Go!")
	fmt.Printf("\n\n%v\nexample of result: %v", trim.String(funcExampleWithVariable), funcVariableResult)
}

// sum takes two integers, num1 and num2, and returns their sum.
// It adds the two input integers and returns the result.
// The function has a single return value.
func sum(num1, num2 int) int {
	return num1 + num2
}

// sumAndsub takes two integers as input and returns two integers.
// The first returned integer is the sum of the input integers,
// and the second returned integer is the difference (num1 - num2) of the input integers.
func sumAndsub(num1, num2 int) (int, int) {
	return num1 + num2, num1 - num2
}

// functionVariable is an anonymous function assigned to a variable.
// It takes a string parameter 'text' and returns the same string.
// This can be used to demonstrate how functions can be assigned to variables in Go.
var functionVariable = func(text string) string {
	return text
}
