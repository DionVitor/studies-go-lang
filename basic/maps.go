package main

import "fmt"

type Location struct {
	Lat, Long float64
}

func main() {
	m := make(map[string]Location)
	m["Bell Labs"] = Location{40.68433, -74.39967}
	fmt.Println(m["Bell Labs"].Long)

	// Maps literals
	var places = map[string]Location{
		"Bell Labs": Location{40.68433, -74.39967},
		"Google": {37.42202, -122.08408}, // Location can be omitted
	}
	fmt.Println(places)

	// Delete item in map
	delete(places, "Bell Labs")
	fmt.Println(places)

	// Key exists?
	key := "Google"
	_, exists := places[key]  // _ is the value in this key
	fmt.Printf("Key: %s / Exists: %t", key, exists)
}
