// InternalArrays demonstrates the performance difference between creating a
// dynamic slice and a pre-allocated slice in Go. It measures the time taken
// to append elements to each type of slice and prints the duration, length,
// and capacity of the slices.
//
// The function performs the following steps:
//  1. Creates a dynamic slice and appends `numElements` elements to it,
//     measuring the time taken.
//  2. Prints the time taken, length, and capacity of the dynamic slice.
//  3. Creates a pre-allocated slice with a capacity of `numElements` and
//     appends `numElements` elements to it, measuring the time taken.
//  4. Prints the time taken, length, and capacity of the pre-allocated slice.
package fundamentals_of_language

import (
	"fmt"
	"time"
)

// numElements defines the number of elements to be used in an array or slice.
// It is set to 1,000,000 for performance testing or large data handling scenarios.
const (
	numElements = 1_000_000
)

var (
	// InternalArraysExample demonstrates the performance difference between creating a
	// dynamic slice and a pre-allocated slice in Go. It measures the time taken
	// to append elements to each type of slice and prints the duration, length,
	// and capacity of the slices.
	InternalArraysExample = `
// InternalArrays demonstrates the performance difference between creating a
// dynamic slice and a pre-allocated slice in Go. It measures the time taken
// to append elements to each type of slice and prints the duration, length,
// and capacity of the slices.

start := time.Now()
var dynamicSlice []int

for idx := range numElements {
	dynamicSlice = append(dynamicSlice, idx)
}

fmt.Println("\nTime to create dynamic slice: ", time.Since(start))
fmt.Println("Length: ", len(dynamicSlice))
fmt.Println("Capacity: ", cap(dynamicSlice))

start = time.Now()
preAllocatedSlice := make([]int, 0, numElements)
for idx := range numElements {
	preAllocatedSlice = append(preAllocatedSlice, idx)
}`
)

// InternalArrays demonstrates the difference in performance between creating a
// dynamic slice and a pre-allocated slice. It measures and prints the time taken
// to create each type of slice, as well as their lengths and capacities. The
// function uses a loop to append elements to both slices and compares the
// efficiency of dynamic allocation versus pre-allocation.
func InternalArrays() {
	fmt.Println(InternalArraysExample)

	start := time.Now()
	var dynamicSlice []int

	for idx := range numElements {
		dynamicSlice = append(dynamicSlice, idx)
	}

	fmt.Println("\nTime to create dynamic slice: ", time.Since(start))
	fmt.Println("Length: ", len(dynamicSlice))
	fmt.Println("Capacity: ", cap(dynamicSlice))

	start = time.Now()
	preAllocatedSlice := make([]int, 0, numElements)
	for idx := range numElements {
		preAllocatedSlice = append(preAllocatedSlice, idx)
	}

	fmt.Println("Time to create pre-allocated slice: ", time.Since(start))
	fmt.Println("Length: ", len(preAllocatedSlice))
	fmt.Println("Capacity: ", cap(preAllocatedSlice))
}
