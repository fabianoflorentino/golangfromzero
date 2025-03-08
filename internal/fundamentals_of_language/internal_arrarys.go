package fundamentals_of_language

import (
	"fmt"
	"time"
)

const (
	numElements = 1_000_000
)

func InternalArrays() {
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
