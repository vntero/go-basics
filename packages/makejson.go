package packages

import (
	"encoding/json"
	"fmt"
)

func MakeJson() {
	// create a map
	person := make(map[string]string)

	// prompts user to enter first name and store it
	fmt.Print("Enter your first name: ")
	var name string
	fmt.Scan(&name)

	// prompts user to enter address and store it
	fmt.Print("Enter your address: ")
	var address string
	fmt.Scan(&address)

	// atrribute name and address to the map
	person["name"] = name
	person["address"] = address

	// use the Marshall() to create JSON object from the map
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Unable to marshall JSON: ", err)
		return
	}

	// print the JSON object
	fmt.Println("Here is the JSON representation of the entered data: ", string(jsonData))
}