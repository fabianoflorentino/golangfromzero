// Package fundamentals_of_language demonstrates fundamental concepts of the Go programming language.
// This package includes examples of struct composition, which is often used as an alternative to inheritance in Go.
// The provided code defines basic structs such as `person`, `employee`, and `student`, and shows how to embed structs
// to reuse fields and methods. The `InheritanceButNot` function illustrates the concept by creating instances of
// `employee` and `student` structs and printing their details.
package fundamentals_of_language

import "github.com/fabianoflorentino/golangfromzero/pkg/trim"

// personStruct is a string containing the definition of the person struct.
// The person struct represents an individual with basic attributes such as name, age, and height.
var (
	personStruct = `
// person represents an individual with basic attributes such as name, age, and height.
type person struct {
	name   string
	age    int
	height int
}
`

	employeeStruct = `
// employee represents an individual who works for a company.
// It embeds the person struct to inherit its fields and methods,
// and adds a company field to store the name of the company the employee works for.
type employee struct {
	person
	company string
}
`

	studentStruct = `
// student represents a student which embeds a person struct and includes additional
// information about the school the student attends.
type student struct {
	person
	school string
}
`
)

// person represents an individual with basic attributes such as name, age, and height.
type person struct {
	name   string
	age    int
	height int
}

// employee represents an individual who works for a company.
// It embeds the person struct to inherit its fields and methods,
// and adds a company field to store the name of the company the employee works for.
type employee struct {
	person
	company string
}

// student represents a student which embeds a person struct and includes additional
// information about the school the student attends.
type student struct {
	person
	school string
}

// InheritanceButNot demonstrates the concept of composition in Go, which is often used as an alternative to inheritance.
// It creates instances of `employee` and `student` structs, both of which embed a `person` struct to reuse its fields.
// The function then prints the details of these instances, showing how the embedded fields can be accessed directly.
func InheritanceButNot() {

	trim := trim.New()

	employee := employee{person: person{name: "John", age: 30, height: 180}, company: "Google"}
	student := student{person: person{name: "Alice", age: 20, height: 160}, school: "MIT"}

	println("\n", trim.String(personStruct))
	println("\n", trim.String(employeeStruct))
	println("\n", trim.String(studentStruct), "\n")

	println("Employee: ", employee.name, employee.age, employee.height, employee.company)
	println("Student: ", student.name, student.age, student.height, student.school)
}
