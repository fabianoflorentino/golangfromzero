package fundamentals_of_language

import (
	"fmt"

	"github.com/fabianoflorentino/golangfromzero/pkg/trim"
)

var (
	exampleArrayWithLength = `
// arrayWithLength demonstrates the creation and initialization of a fixed-length array in Go.
// It declares an array of 5 strings, assigns a value to the first element, and prints it.

Example 1:

var arr1 [5]string
arr1[0] = "a"

Example 2:

arr2 := [5]string{"a", "b", "c", "d", "e"}
`
)

func ArraysAndSlices() {
	trim := trim.New()

	println("\nWith fixed-length arrays:\n")
	println(trim.String(exampleArrayWithLength))
	arrayWithLength()
}

// arrayWithLength demonstrates the creation and initialization of a fixed-length array in Go.
// It declares an array of 5 strings, assigns a value to the first element, and prints it.
func arrayWithLength() {
	var arr1 [5]string
	arr1[0] = "a"
	println("arr1[0]: ", arr1[0])

	arr2 := [5]string{"a", "b", "c", "d", "e"}
	fmt.Println(arr2)
}
