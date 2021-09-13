package main

import (
	"fmt"
)

// Maps - similar to dictionaries of hashtables with a characteristic key/value pair
// Major constraint of maps is their keys and values must have the same data-type defined in the declaration
func main() {
	stateCovidCases := map[string]int{
		"Lagos":  23000,
		"Osun":   12000,
		"Taraba": 15000,
		"Abuja":  12e3,
		"Ogun":   21000,
	}

	fmt.Printf("%v, %T\n", stateCovidCases, stateCovidCases)
	// Members of a map are accessed via their keys
	fmt.Printf("%v, %T\n", stateCovidCases["Lagos"], stateCovidCases["Lagos"])
	// Manipulating values in a map
	stateCovidCases["Lagos"] = 340000
	fmt.Printf("%v, %T\n", stateCovidCases["Lagos"], stateCovidCases["Lagos"])

	// Remove a value from a map based on its key
	delete(stateCovidCases, "Ogun")
	fmt.Println(stateCovidCases)
	// Length of the keys in the map
	fmt.Println(len(stateCovidCases))

	// map1 := make(map[string]int, 10)
	// map1["key"] = 20
	// fmt.Println(len(map1))
}
