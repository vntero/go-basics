package packages

import (
	"fmt"
	"strings"
)

func Findian() {
	var input string 
	fmt.Print("Enter a string: ")
	fmt.Scan(&input)

	var loweredInput string = strings.ToLower(input)

	var startsWithI bool = strings.HasPrefix(loweredInput, "i")
	var containsA bool = strings.Contains(loweredInput, "a")
	var endsWithN bool = strings.HasSuffix(loweredInput, "n")

	var isValid bool = startsWithI && containsA && endsWithN

	if isValid {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not found!")
	}
}