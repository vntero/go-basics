package packages

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// make a type called animal
type Animal struct {
	food       string
	locomotion string
	noise      string
}

// make three methods
func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.noise)
}

func SelectAnimal() {
	// predefine animals
	cow := Animal{
		food:       "grass",
		locomotion: "walk",
		noise:      "moo",
	}

	bird := Animal{
		food:       "worms",
		locomotion: "fly",
		noise:      "peep",
	}

	snake := Animal{
		food:       "mice",
		locomotion: "slither",
		noise:      "hsss",
	}

	// create a map to link animalName with corresponding type
	animals := map[string]Animal{
		"cow":   cow,
		"bird":  bird,
		"snake": snake,
	}

	for {
		// prompt user
		fmt.Println(">")

		// grab what user typed
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		elements := strings.Fields(input)

		// check correct input
		if len(elements) != 2 {
			fmt.Println("Invalid input. Please provide 1 animal and 1 action.")
		}

		animalName, animalAction := elements[0], elements[1]
		animal, ok := animals[animalName]

		// check correct animal name
		if !ok {
			fmt.Println("Enter a valid animal.")
			continue
		}

		switch animalAction {
		case "eat":
			animal.Eat()
		case "move":
			animal.Move()
		case "speak":
			animal.Speak()
		default:
			fmt.Println("Enter a valid action.")
		}
	}
	
}
