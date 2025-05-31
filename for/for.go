package main

import "fmt"

func rangeIterator(start, end int) []int {
	size := end - start
	if size <= 0 {
		return []int{}
	}

	rangeSlice := make([]int, size)
	for i := range rangeSlice {
		rangeSlice[i] = start + i
	}

	return rangeSlice

}

func main() {
	myRange := rangeIterator(0, 5)

	fmt.Println("Object being iterated:", myRange)

	// Basic structure
	fmt.Println("Using basic for structure")
	for i := 0; i < len(myRange); i++ {
		fmt.Println(myRange[i])
	}

	fmt.Println()

	// Range structure (more modern)
	fmt.Println("Using range structure in for loop")
	for _, i := range myRange {
		fmt.Println(i)
	}
}