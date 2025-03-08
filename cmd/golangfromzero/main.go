package main

import (
	"fmt"

	"github.com/fabianoflorentino/golangfromzero/internal/fundamentals_of_language"
)

func main() {
	fmt.Println("Leard Golang from Zero!")

	fmt.Println("\n5. External Packages\n")
	fmt.Println("fabianoflorentino@outlook.com: ", fundamentals_of_language.ExternalPackages("fabianoflorentino@outlook.com"))
	fmt.Println("invalid: ", fundamentals_of_language.ExternalPackages("invalid"))

	fmt.Println("\n6. Variables")
	fmt.Println(fundamentals_of_language.TypeVariables())

	fmt.Println("\n7. Type of Data")
	fmt.Println(fundamentals_of_language.TypeOfData())

	fmt.Println("\n8. Type of Functions")
	fundamentals_of_language.TypeOfFunctions()

	fmt.Println("\n9. Operators")
	fundamentals_of_language.Operators()

	fmt.Println("\n10. Type of Struct")
	fundamentals_of_language.TypeOfStruct()

	fmt.Println("\n11. Inheritance but not")
	fundamentals_of_language.InheritanceButNot()

	fmt.Println("\n12. Arrays and Slices")
	fundamentals_of_language.ArraysAndSlices()

	fmt.Println("\n13. Pointers")
	fundamentals_of_language.Pointers()

	fmt.Println("\n14. Internal Arrays")
	fundamentals_of_language.InternalArrays()

	fmt.Println("\n15. Maps")
	fundamentals_of_language.Maps()
}
