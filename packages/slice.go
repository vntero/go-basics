package packages

import (
	"fmt"
	"sort"
	"strconv"
)

// 4. Sorted slice
func SortedSlice() {
	//create an empty integer slice of length 3
	slice := make([]int, 0, 3)

	// program should be written as a loop
	for {
		// prompt the user
		fmt.Print("Enter an integer to be added to the slice or 'X' to close program: ")
		var input string
		fmt.Scan(&input)

		// quit when the user enters the character X
		if input == "X" || input == "x" {
			fmt.Println("Program closed.")
			break
		}

		// convert input from string to int
		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("The input entered is not valid. Try again.")
		}

		// add the integer to the slice
		slice = append(slice, num)

		// sort the slice
		sort.Ints(slice)

		// prints the content of the slice
		fmt.Println(slice)
	}
}