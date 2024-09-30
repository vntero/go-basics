package packages

import (
	"fmt"
	"math"
)

// 2. Truncate floating point number
func TruncateFloat() {
	var input float64

	fmt.Println("Enter a floating-point number:")
	fmt.Scan(&input)

	truncated := math.Trunc(input)

	fmt.Println("The truncated version of the number", input, "is", truncated)
}