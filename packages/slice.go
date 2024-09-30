package packages

import "fmt"

func SortedSlice() {
	//create an empty integer slice of length 3
	slice := make([]int, 0, 3)

	// program should be written as a loop
	for {
		// prompt the user
		fmt.Print("Enter an integer to be added to the slice or 'X' to close program: ")
		var input string
		fmt.Scan(&input)

		// Quit when the user enters the character X
		if input == 'X' || input == 'x' {
			fmt.Print("Program closed.")
		}
	}
}