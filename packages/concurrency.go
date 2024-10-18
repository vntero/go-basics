package packages

import "fmt"

// partition an array into 4 parts
// each of which is sorted by different goroutine
// each partition should be approximately equal size
// main goroutine should merge the 4 sorted subarrays into one large sorted array

// each goroutine which sorts a quarter of the array should print the subarray that it will sort
// when sorting is complete, the main goroutine should print the entire sorted list

func SortArray() {
	// prompt the user to input a series of integers
	fmt.Println("Enter a series of integers:")
}
