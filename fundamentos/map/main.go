package main

import (
	"fmt"

	"github.com/douglasandradeee/GoExercicies/advancedStructs/model"
)

var cache map[string]model.Immobile

func main() {

	cache = make(map[string]model.Immobile)

	house := model.Immobile{}
	house.Name = "Yellow House"
	house.X = 18
	house.Y = 25
	house.SetValue(60000)

	cache[house.Name] = house

	fmt.Println("There is a yellow house in the cache?")
	immobile, found := cache["Yellow House"]
	if found == true {
		fmt.Printf("Yes, and what I found was: %+v\r\n", immobile)
	}

	apartment := model.Immobile{}
	apartment.Name = "Blue Apartment"
	apartment.X = 19
	apartment.Y = 26
	apartment.SetValue(70000)

	// Add items in the map
	cache[apartment.Name] = apartment

	// Give us the size of the map.
	fmt.Println("How many items there is in the cache? ", len(cache))

	// List map Items.
	for key, immobile := range cache {
		fmt.Printf("key[%s] =  %+v\r\n", key, immobile)
	}

	// Find and Delete items in the map.
	immobile, found = cache["Yellow House"]
	if found == true {
		delete(cache, immobile.Name)
	}
	fmt.Println("How many items there is in the cache? ", len(cache))
}
