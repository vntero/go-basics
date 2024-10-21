package packages

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func sorter(a []int, wg *sync.WaitGroup) {
	defer wg.Done()

	// EACH GOROUTINE WHICH SORTS A QUARTER OF THE ARRAY SHOULD PRINT THE SUB-ARRAY THAT IT WILL SORT
	fmt.Println("Starting to sort the following sub-array:", a)

	sort.Ints(a)

	// print again once sub-array is sorted
	fmt.Println("Sorted sub-array:", a)
}

func merger(a1, a2, a3, a4 []int) []int  {
	a := append(a1, a2...)
	a = append(a, a3...)
	a = append(a, a4...)
	sort.Ints(a)
	return a
}

func SortArray() {
	var wg sync.WaitGroup

	// PROMPT THE USER TO INPUT A SERIES OF INTEGERS
	fmt.Println("Enter a series of at least 12 integers separated by space:")

	// grab what user entered
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	parts := strings.Fields(input)

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

	// expect at least 12 integers
	if len(numbers) < 12 {
		fmt.Println("Insufficient number of integers entered. Try again.")
		return
	}

	// PARTITION ARRAY INTO 4 PARTS
	// EACH PARTITION SHOULD BE APPROX. EQUAL SIZE
	n := len(numbers) / 4
	array1 := numbers[0:n]
	array2 := numbers[n : 2*n]
	array3 := numbers[2*n : 3*n]
	array4 := numbers[3*n:]

	// EACH OF WHICH IS SORTED BY DIFFERENT GOROUTINE
	wg.Add(4)
	go sorter(array1, &wg)
	go sorter(array2, &wg)
	go sorter(array3, &wg)
	go sorter(array4, &wg)
	wg.Wait()

	// main goroutine should merge the 4 sorted subarrays into one large sorted array
	result := merger(array1, array2, array3, array4)

	// when sorting is complete, the main goroutine should print the entire sorted list
	fmt.Println("Here's the final result:", result)
}
