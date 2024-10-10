package packages

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Animal interface with three methods
type Animal interface {
	Eat()
	Move()
	Speak()
}

// Cow type and its methods
type Cow struct{}

func (c Cow) Eat()   { fmt.Println("grass") }
func (c Cow) Move()  { fmt.Println("walk") }
func (c Cow) Speak() { fmt.Println("moo") }

// Bird type and its methods
type Bird struct{}

func (b Bird) Eat()   { fmt.Println("worms") }
func (b Bird) Move()  { fmt.Println("fly") }
func (b Bird) Speak() { fmt.Println("peep") }

// Snake type and its methods
type Snake struct{}

func (s Snake) Eat()   { fmt.Println("mice") }
func (s Snake) Move()  { fmt.Println("slither") }
func (s Snake) Speak() { fmt.Println("hsss") }

func createAnimal() {
	// Map to store created animals by name
	animals := make(map[string]Animal)

	// Create a scanner to read user input
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")

		// Grab user input
		scanner.Scan()
		input := scanner.Text()

		// Split input into parts
		parts := strings.Fields(input)

		// Check if input has exactly 3 parts
		if len(parts) != 3 {
			fmt.Println("Invalid input. Please use 'newanimal' or 'query' followed by 2 parameters.")
			continue
		}

		command := parts[0] // newanimal or query
		name := parts[1]    // animal name
		param := parts[2]   // animal type or action

		switch command {
		case "newanimal":
			// Create a new animal based on the provided type (cow, bird, snake)
			var animal Animal

			switch param {
			case "cow":
				animal = Cow{}
			case "bird":
				animal = Bird{}
			case "snake":
				animal = Snake{}
			default:
				fmt.Println("Invalid animal type. Please choose 'cow', 'bird', or 'snake'.")
				continue
			}

			animals[name] = animal
			fmt.Println("Created it!")

		case "query":
			// Query an animal's behavior (eat, move, speak)
			animal, exists := animals[name]
			if !exists {
				fmt.Println("Animal not found.")
				continue
			}

			switch param {
			case "eat":
				animal.Eat()
			case "move":
				animal.Move()
			case "speak":
				animal.Speak()
			default:
				fmt.Println("Invalid query. Use 'eat', 'move', or 'speak'.")
			}

		default:
			fmt.Println("Invalid command. Use 'newanimal' to create or 'query' to ask about an animal.")
		}
	}
}
