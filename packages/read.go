package packages

import "fmt"

func Read() {
		// reads information from a file

	// represents it in a slice of structs

	// prompt the user for the name of the text file
	fmt.Print("Enter the name of text file: ")

	// successively read each line of the text file

	// create a struct which contains the first and last names found in the file
	type Name struct {
		fname string
		lname string
	}

	// each struct created will be added to the file

	// iterate through the slice of structs and print the first and last names found in each struct
}