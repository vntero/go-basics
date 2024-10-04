package packages

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Swap(slice []int, i int) {
	slice[i], slice[i+1] = slice[i+1], slice[i]
}

func BubbleSort(slice []int) {
	n := len(slice)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if slice[j] > slice[j+1] {
				Swap(slice, j)
			}
		}
	}
}

func BubbleSorted() {
	// prompt the user
	fmt.Println("Type in a sequence up to 10 integers separated by spaces.")

	// grab what user typed
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("raw input", input)

	// split input and check if length is 10
	elements := strings.Fields(input)
	fmt.Println("clean input", elements)

	// convert input from string to integers
	numbers := make([]int, 0, len(elements))
	for _, elem := range elements {
		num, err := strconv.Atoi(elem) // convert each string element to an integer
		if err != nil {
			fmt.Println("Invalid input. Please enter valid integers.")
			return
		}
		numbers = append(numbers, num)
	}

	fmt.Println("converted input", numbers)

	// sort the numbers using BubbleSort
	BubbleSort(numbers)

	// print the sorted result
	fmt.Println("sorted result", numbers)
}
