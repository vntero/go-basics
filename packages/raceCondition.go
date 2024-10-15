package packages

import (
	"fmt"
	"sync"
)

// shared global variable
var counter int

// function that manipulayes the global variable
func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		counter = counter + 1
	}
}

func RaceCondition() {
	var wg sync.WaitGroup

	wg.Add(2)
	go increment(&wg)
	go increment(&wg)

	wg.Wait()

	fmt.Println("Final counter value:", counter)
}

