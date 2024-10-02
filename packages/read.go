package packages

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// struct containing first and last names found in file
type Name struct {
	fname string
	lname string
}

// variable to store names found in file
var names []Name

func Read() {
	// prompt user for file name
	fmt.Println("Enter name of file you'd like to open:")
	var fileName string
	fmt.Scan(&fileName)

	// open file
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Unable to open file: ", err)
		return
	}
	defer file.Close()

	// read each line of the text file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")

		if len(parts) != 2 {
			fmt.Print("Wrong name format. Please use only first and last name.")
			continue
		}

		// store each name in global variable
		name := Name{
			fname: parts[0],
			lname: parts[1],
		}
		names = append(names, name)

		if err := scanner.Err(); err != nil {
			fmt.Println("Unable to read file: ", err)
			return
		}
	}

	// iterate through the slice of structs and print the first and last names found in each struct
	fmt.Println("\nNames found in the file: ")
	for _, name := range names {
		fmt.Println("First name:", name.fname, "- Last name:", name.lname)
	}
}