package packages

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

// main goroutine should merge the 4 sorted subarrays into one large sorted array

// each goroutine which sorts a quarter of the array should print the subarray that it will sort
// when sorting is complete, the main goroutine should print the entire sorted list
func sorter(a []int, wg *sync.WaitGroup) {
	defer wg.Done()

	n := len(a)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if a[j] > a[j+1] {
				Swap(a, j)
			}
		}
	}
}

func SortArray() {
	var wg sync.WaitGroup

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
	n := len(numbers) / 4
	array1 := numbers[0:n]
	array2 := numbers[n : 2*n]
	array3 := numbers[2*n : 3*n]
	array4 := numbers[3*n:]

	// handle spare parts
	// EACH PARTITION SHOULD BE APPROX. EQUAL SIZE
	if len(numbers)%4 != 0 {
		array4 = append(array4, numbers[4*n:]...)
	}

	// EACH OF WHICH IS SORTED BY DIFFERENT GOROUTINE
	wg.Add(4)
	go sorter(array1, &wg)
	go sorter(array2, &wg)
	go sorter(array3, &wg)
	go sorter(array4, &wg)
	wg.Wait()

	log.Printf("array1: %v\n", array1)
	log.Printf("array2: %v\n", array2)
	log.Printf("array3: %v\n", array3)
	log.Printf("array4: %v\n", array4)
}
