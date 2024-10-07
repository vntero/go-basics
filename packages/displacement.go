package packages

import (
	"fmt"
)

func GenDisplaceFn(a, v, s float64) func(float64) float64 {
	return func(t float64) float64 {
		return 0.5*a*t*t + v*t + s
	}
}


func Displacement() {
	var a, vo, so, t float64

	// prompt user for initial values
	fmt.Print("Enter acceleration (a): ")
	fmt.Scan(&a)

	fmt.Print("Enter initial velocity (vo): ")
	fmt.Scan(&vo)

	fmt.Print("Enter initial displacement (so): ")
	fmt.Scan(&so)

	// generate displacement function
	fn := GenDisplaceFn(a, vo, so)

	// prompt user for time
	fmt.Print("Enter time (t): ")
	fmt.Scan(&t)

	// compute displacement
	displacement := fn(t)
	fmt.Println("Displacement after", t, "seconds is", displacement)
}

// // prompt the user for values
// fmt.Println("Enter values for acceleration, initial velocity, and initial displacement.")

// // grab what the user typed
// reader := bufio.NewReader(os.Stdin)
// input, _ := reader.ReadString('\n')

// // grab each value individually
// elements := strings.Fields(input)

// // convert values from string to int
// numbers := make([]int, 0, len(elements))
// for _, element := range elements {
// 	num, err := strconv.Atoi(element)
// 	if err != nil {
// 		fmt.Println("Invalid input. Please enter valid integers.")
// 		return
// 	}
// 	numbers = append(numbers, num)
// }

// // prompt user for time
// fmt.Println("Enter value for time.")