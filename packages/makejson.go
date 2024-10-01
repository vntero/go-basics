package packages

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func MakeJson() {
		// create a map
		person := make(map[string]string)

		// prompt user to enter first name and store it
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter your first name: ")
		name, _ :=  reader.ReadString('\n')
	
		// prompt user to enter address and store it
		fmt.Print("Enter your address: ")
		address, _ := reader.ReadString('\n')
	
		// atrribute name and address to the map
		person["name"] = name[:len(name)-1]
		person["address"] = address[:len(address)-1]
	
		// use the Marshall() to create JSON object from the map
		jsonData, err := json.Marshal(person)
		if err != nil {
			fmt.Println("Unable to marshall JSON: ", err)
			return
		}
	
		// print the JSON object
		fmt.Println("Here is the JSON representation of the entered data: ", string(jsonData))
}