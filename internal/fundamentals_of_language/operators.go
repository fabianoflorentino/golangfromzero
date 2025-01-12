// Package fundamentals_of_language demonstrates basic concepts and operations
// in the Go programming language. This package includes examples of various
// operators, control structures, and other fundamental language features.
package fundamentals_of_language

import "github.com/fabianoflorentino/golangfromzero/pkg/trim"

var (
	// assignmentOperatorsExample demonstrates the use of assignment operators in Go.
	// It includes examples of:
	// - = (simple assignment) for assigning a value to a variable that has already been declared.
	// - := (short variable declaration) for declaring and initializing a variable in one step.
	assignmentOperatorsExample = `
// The following code demonstrates the use of assignment operators in Go.
// - = (simple assignment)
// - := (short variable declaration)
var var1 string = "Hello"
var2 := "World"
`

	// refactTernaryOperator demonstrates how to achieve ternary operator-like functionality in Go
	// using an if-else statement. Go does not have a built-in ternary operator, so conditional
	// assignments must be done using standard if-else constructs.
	refactTernaryOperator = `
// The ternary operator is not available in Go.
// However, you can use an if-else statement to achieve similar functionality.
// Example:
if condition {
	value = trueValue
} else {
	value = falseValue
}
`
)

// Operators demonstrates the use of arithmetic operators in Go.
// It performs addition, subtraction, multiplication, division, and modulus operations
// on two integers (num1 and num2) and prints the results.
func Operators() {

	trim := trim.New()

	// The following code demonstrates the use of arithmetic operators on two integers.
	// The result of the addition operation is of type int.
	// The result of the subtraction operation is of type int.
	// The result of the division operation is an integer division.
	// The result of the modulus operation is the remainder of the division.
	num1, num2 := 10, 20
	sum, sub, mul, div, mod := arithmeticOperators(num1, num2)
	println("\nSum: 10 + 20", sum, "Sub: 10 - 20", sub, "Mul: 10 * 20", mul, "Div: 10 / 20", div, "Mod: 10 % 20", mod)

	// The following code demonstrates the use of arithmetic operators on two integers of the same type.
	// The result of the addition operation is of the same type as the operands.
	// The result of the addition operation is of type int16.
	num1Int16, num2Int16 := int16(10), int16(20)
	sumInt16 := arithmeticOperatorsSameType(num1Int16, num2Int16)
	println("\nSum (int16):", sumInt16, "does not works if you use different types of data, Example: int16 + int32\n")

	// The following code demonstrates the use of assignment operators in Go.
	println(trim.String(assignmentOperatorsExample))
	assignmentOperators()

	// The following code demonstrates the use of relational operators in Go.
	relationOperators()

	// The following code demonstrates the use of logical operators in Go.
	logicalOperators()

	// The following code demonstrates the use of unary operators in Go.
	unaryOperators()

	// The following code demonstrates the use of the ternary operator in Go.
	println("\nGo does not implemtnet the ternary operator (?:) like C, C++, Java, and other languages.\n")
	println(trim.String(refactTernaryOperator))

}

// arithmeticOperators performs basic arithmetic operations on two integers.
// It returns the results of addition, subtraction, multiplication, division, and modulus operations.
// Parameters:
// - num1: The first integer operand.
// - num2: The second integer operand.
// Returns:
// - The result of num1 + num2 (addition).
// - The result of num1 - num2 (subtraction).
// - The result of num1 * num2 (multiplication).
// - The result of num1 / num2 (integer division).
// - The result of num1 % num2 (modulus).
func arithmeticOperators(num1, num2 int) (int, int, int, int, int) {
	return num1 + num2, num1 - num2, num1 * num2, num1 / num2, num1 % num2
}

// arithmeticOperatorsSameType takes two int16 numbers as input and returns their sum.
// This function demonstrates the use of arithmetic addition operator with operands of the same type.
func arithmeticOperatorsSameType(num1, num2 int16) int16 {
	return num1 + num2
}

// assignmentOperators demonstrates the use of assignment operators in Go.
// It initializes two string variables, var1 and var2, with the values "Hello" and "World" respectively,
// and then prints them to the console.
func assignmentOperators() {
	var var1 string = "Hello"
	var2 := "World"

	println("\nAssignment Operators: ", var1, var2)
}

// relationOperators demonstrates the use of relational operators in Go.
// It compares two integers, num1 and num2, and prints the results of
// various relational operations (==, !=, >, <, >=, <=) between them.
func relationOperators() {
	num1, num2 := 10, 20
	println("\nRelational Operators:\n")
	println("10 == 20", num1 == num2)
	println("10 != 20", num1 != num2)
	println("10 > 20", num1 > num2)
	println("10 < 20", num1 < num2)
	println("10 >= 20", num1 >= num2)
	println("10 <= 20", num1 <= num2)
}

// logicalOperators demonstrates the use of logical operators in Go.
// It prints the results of logical AND, OR, and NOT operations
// using predefined boolean values.
func logicalOperators() {
	trueValue, falseValue := true, false

	println("\nLogical Operators:\n")
	println("true && false", trueValue && falseValue) // Logical AND
	println("true || false", trueValue || falseValue) // Logical OR
	println("!true", !trueValue)                      // Logical NOT
}

// unaryOperators demonstrates the use of unary operators in Go.
// It performs various unary operations on an integer (num) and a boolean value (true).
// The operations include:
// - + (positive sign)
// - - (negative sign)
// - ^ (bitwise NOT)
// - ++ (increment)
// - -- (decrement)
func unaryOperators() {
	num := 10

	println("\nUnary Operators:\n")
	println("+10", +num) // Positive
	println("-10", -num) // Negative
	println("^10", ^num) // Bitwise NOT

	num++
	println("num++:", num) // Increment

	num = 10
	num--
	println("num--:", num) // Decrement
}
