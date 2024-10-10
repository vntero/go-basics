package packages

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// animal interface
type Animal interface {
	Eat()
	Move()
	Speak()
}

// Cow struct and its methods
type Cow struct{}

func (c Cow) Eat()   { fmt.Println("grass") }
func (c Cow) Move()  { fmt.Println("walk") }
func (c Cow) Speak() { fmt.Println("moo") }

// Bird struct and its methods
type Bird struct{}

func (b Bird) Eat()   { fmt.Println("worms") }
func (b Bird) Move()  { fmt.Println("fly") }
func (b Bird) Speak() { fmt.Println("peep") }

// Snake struct and its methods
type Snake struct{}

func (s Snake) Eat()   { fmt.Println("mice") }
func (s Snake) Move()  { fmt.Println("slither") }
func (s Snake) Speak() { fmt.Println("hsss") }

func CreateAnimal() {
	// store the created animals
	animals := make(map[string]Animal)

	// grab user input
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")

		// grab user input
		scanner.Scan()
		input := scanner.Text()

		// split input
		parts := strings.Fields(input)

		// check input
		if len(parts) != 3 {
			fmt.Println("Invalid command. Please enter 'newanimal' or 'query' followed by 2 parameters.")
			continue
		}

		command := parts[0]
		name := parts[1]
		param := parts[2]

		// newanimal animalName animalType
		if command == "newanimal" {
			var animal Animal

			switch param {
			case "cow":
				animal = Cow{}
			case "bird":
				animal = Bird{}
			case "snake":
				animal = Snake{}
			default:
				fmt.Println("Invalid animal type. Choose 'cow', 'bird', or 'snake'.")
				continue
			}

			animals[name] = animal
			fmt.Println("Created it!")

		} else if command == "query" {
			// query animalName animalAction
			animal, exists := animals[name]
			if !exists {
				fmt.Println("No such animal found.")
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
		} else {
			// handle invalid commands
			fmt.Println("Invalid command. Use 'newanimal' to create an animal or 'query' to ask about an animal.")
		}
	}
}
