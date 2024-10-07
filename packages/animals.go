package packages

import "fmt"

// make a type called animal
type Animal struct {
	food       string
	locomotion string
	noise      string
}

// make three methods
func (a Animal) Eat(){
	fmt.Println(a.food)
}

func (a Animal) Move(){
	fmt.Println(a.locomotion)
}

func (a Animal) Speak(){
	fmt.Println(a.noise)
}



func Animals() {
	// predefined animals
	cow := Animal{
		food: "grass",
		locomotion: "walk",
		noise: "moo",
	}

	bird := Animal{
		food: "worms",
		locomotion: "fly",
		noise: "peep",
	}

	snake := Animal{
		food: "mice",
		locomotion: "slither",
		noise: "hsss",
	}


	// prompt user
	fmt.Println("Enter the name of the animal followed by name of information requested.")

	fmt.Println(cow)
}
