package packages

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// each of which is sorted by different goroutine
// each partition should be approximately equal size
// main goroutine should merge the 4 sorted subarrays into one large sorted array

// each goroutine which sorts a quarter of the array should print the subarray that it will sort
// when sorting is complete, the main goroutine should print the entire sorted list

func SortArray() {
	// PROMPT THE USER TO INPUT A SERIES OF INTEGERS
	fmt.Println("Enter a series of at least 12 integers separated by space:")

	// grab what user entered
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	log.Printf("input: %v\n", input)
	parts := strings.Fields(input)
	log.Printf("parts: %v\n", parts)

	// expect at least 12 integers
	if len(input) < 12 {
		fmt.Println("Insufficient numbers of integers entered. Try again.")
		return
	}

	// convert array of strings into array of numbers
	numbers := make([]int, 0, len(parts))
	for _, elem := range parts {
		num, err := strconv.Atoi(elem) // convert each string element to an integer
		if err != nil {
			fmt.Println("Invalid input. Please enter valid integers.")
			return
		}
		numbers = append(numbers, num)
	}
	log.Printf("numbers: %v\n", numbers)

	// PARTITION ARRAY INTO 4 PARTS
}
