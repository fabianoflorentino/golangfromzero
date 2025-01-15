package main

import (
	"github.com/fabianoflorentino/golangfromzero/internal/fundamentals_of_language"
)

func main() {
	println("Leard Golang from Zero!")

	println("\n5. External Packages\n")
	println("fabianoflorentino@outlook.com: ", fundamentals_of_language.ExternalPackages("fabianoflorentino@outlook.com"))
	println("invalid: ", fundamentals_of_language.ExternalPackages("invalid"))

	println("\n6. Variables")
	println(fundamentals_of_language.TypeVariables())

	println("\n7. Type of Data")
	println(fundamentals_of_language.TypeOfData())

	println("\n8. Type of Functions")
	fundamentals_of_language.TypeOfFunctions()

	println("\n9. Operators")
	fundamentals_of_language.Operators()

	println("\n10. Type of Struct")
	fundamentals_of_language.TypeOfStruct()

	println("\n11. Inheritance but not")
	fundamentals_of_language.InheritanceButNot()

	println("\n12. Arrays and Slices")
	fundamentals_of_language.ArraysAndSlices()

	println("\n13. Pointers")
	fundamentals_of_language.Pointers()
}
