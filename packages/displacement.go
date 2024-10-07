package packages

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GenDisplaceFn(a float64, v float64, s float64) func(t float64) {

}


func Displacement() {
	// prompt the user for values
	fmt.Println("Enter values for acceleration, initial velocity, and initial displacement.")

	// grab what the user typed
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	// grab each value individually
	elements := strings.Fields(input)

	// convert values from string to int
	numbers := make([]int, 0, len(elements))
	for _, element := range elements {
		num, err := strconv.Atoi(element)
		if err != nil {
			fmt.Println("Invalid input. Please enter valid integers.")
			return
		}
		numbers = append(numbers, num)
	}

	// prompt user for time
	fmt.Println("Enter value for time.")
}
