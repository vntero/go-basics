package packages

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func BubbleSort() {
	// declare an empty slice for storing user input
	// slice := make([]int, 10)

	// prompt the user
	fmt.Println("Type in a sequence up to 10 integers separated by spaces.")

	// grab what user typed
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("raw input", input)

	// split input and check if length is 10
	elements := strings.Fields(input)
	// if len(elements) > 10 {
	// 	fmt.Println("Too many integers entered, please type in up to 10.")
	// 	return
	// }

	fmt.Println("clean input", elements)

	// convert input from string to integers
	numbers := make([]int, 0, len(elements))
	for i := 1; i <= len(elements); i++ {
		numbers = append(numbers, i)
	}

	fmt.Println("converted input", numbers)
}