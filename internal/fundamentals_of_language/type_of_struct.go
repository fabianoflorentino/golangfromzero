package fundamentals_of_language

var (
	structExample = `
// Person is a struct that represents a person.
// It has two fields: name and age.
type Person struct {
	name string
	age  int
}
`

	anonymousFieldExample = `
// address is a struct that represents an address.
// It has two fields: city and state.
type address struct {
	city, state string
}
`

	structExampleWithAnonymousField = `
// Person is a struct that represents a person.
// It has three fields: name, age, and address.
// The address field is an anonymous field of type address.
type Person struct {
	name    string
	age     int
	address // anonymous field
}
`
)

type address struct {
	city, state string
}

type Person struct {
	name    string
	age     int
	address // anonymous field
}

func TypeOfStruct() {
	p1 := Person{name: "Alice", age: 25}
	p2 := Person{name: "Bob", age: 30, address: address{city: "New York", state: "NY"}}

	println(structExample)
	println("Name: ", p1.name, "Age: ", p1.age)

	println(anonymousFieldExample)
	println(structExampleWithAnonymousField)
	println("Name: ", p2.name, "Age: ", p2.age, "City: ", p2.city, "State: ", p2.state)
}
