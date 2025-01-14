// Package fundamentals_of_language provides examples and explanations of fundamental concepts in the Go programming language.
// This package includes demonstrations of variable assignment, the use of pointers, and the differences between using pointers and not using pointers.
package fundamentals_of_language

import (
	"fmt"

	"github.com/fabianoflorentino/golangfromzero/pkg/trim"
)

// exampleWithoutPointer demonstrates the behavior of variables without using pointers.
// In this example, `a` is assigned the value 10, and `b` is assigned the value of `a`.
// When `a` is incremented, `b` remains unchanged because `b` holds a copy of the value of `a`.
var (
	exampleWithoutPointer = `
var a int = 10
var b int = a

fmt.Println(a, b) // 10 10

a++

fmt.Println(a, b) // 11 10
`

	exampleWithPointer = `
var a int = 10
var b *int = &a
var a int = 10
var b *int = &a

fmt.Println(a, *b) // 10 10

a++

fmt.Println(a, *b) // 11 11
`
)

// Pointers demonstrates the difference between using pointers and not using pointers in Go.
// It prints the results of two examples: one without using pointers and one using pointers.
func Pointers() {
	trim := trim.New()

	println("\nWithout pointers:\n")
	println(trim.String(exampleWithoutPointer))
	withoutPointer()

	println("\nUsing pointers:\n")
	println(trim.String(exampleWithPointer))
	withPointer()
}

// withoutPointer demonstrates how variable assignment works in Go when not using pointers.
// It shows that assigning one variable to another creates a copy, so changes to the original
// variable do not affect the copied variable.
func withoutPointer() {
	var a int = 10
	var b int = a

	fmt.Println(a, b) // 10 10

	a++

	fmt.Println(a, b) // 11 10
}

// withPointer demonstrates the use of pointers in Go.
// It initializes an integer variable 'a' with a value of 10 and a pointer 'b' that points to 'a'.
// It prints the values of 'a' and the value pointed to by 'b', which are initially the same.
// Then, it increments 'a' and prints the values again, showing that the value pointed to by 'b' also changes.
func withPointer() {
	var a int = 10
	var b *int = &a

	fmt.Println(a, *b) // 10 10

	a++

	fmt.Println(a, *b) // 11 11
}
