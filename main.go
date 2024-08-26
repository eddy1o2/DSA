package main

import (
	"DSA/array"
	"fmt"
)

func main() {
	fmt.Println("Pivot Index: ", array.PivotIndex([]int{1, 7, 3, 6, 5, 6}))

	fmt.Println("Remove Element To Make Strictly Increasing: ", array.CanBeIncreasing([]int{1, 2, 10, 5, 7}))

	fmt.Println("Remove Element: ", array.RemoveElement([]int{3, 2, 2, 3}, 3))

	fmt.Println("Remove Duplicate Sorted Array: ", array.RemoveDuplicates([]int{1, 1, 2}))
}
