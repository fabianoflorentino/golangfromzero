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

Example 3:

arr3 := [...]int{1, 2, 3, 4, 5}
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
	println("\nResults:")

	var arr1 [5]string
	arr1[0] = "a"
	println("\nResult Example 1: ", arr1[0])

	arr2 := [5]string{"a", "b", "c", "d", "e"}
	fmt.Println("\nResult Example 2: ", arr2)

	arr3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println("\nResult Example 3: ", arr3)
}
