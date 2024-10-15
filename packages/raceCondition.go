package packages

import (
	"fmt"
	"sync"
)

// shared global variables
var fullNames = []string{}
var firstNames = []string{"Hugo", "Darren", "Marcus", "Cristiano", "Pep", "Jose", "Ryan", "Paul", "Roy", "Gary"}
var	lastNames = []string{"Neville", "Keane", "Scholes", "Giggs", "Mourinho", "Guardiola", "Ronaldo", "Rashford", "Fletcher", "Antero"}

// function that manipulayes the global variable
func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fullNames = append(fullNames, firstNames[i])
		fullNames= append(fullNames, lastNames[i])
	}
}

func RaceCondition() {
	var wg sync.WaitGroup
	

	// start two goroutines
	wg.Add(2)
	go increment(&wg)
	go increment(&wg)

	// wait for both goroutines to finish
	wg.Wait()

	// fullNames slice is being shared between 2 goroutines
	// and both goroutines modify it by appending data to it
	// since both goroutines are modifying the shared slice at the same time without any locks or synchronization mechanisms
	// it may lead to unpredictable behaviour
	// append is not thread-safe, there is a chance that the slice's internal array gets reallocated while both goroutines are writing it
	// that's why when printing the total of full names (combination of firstNames + lastNames), each time we get a different value
	fmt.Println("Number of full names (first + last):", len(fullNames) / 2)
	fmt.Println("Full names generated (combination of firstNames + lastNames):", fullNames)
}

