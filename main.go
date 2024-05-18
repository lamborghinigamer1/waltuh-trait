package main

import (
	"fmt"
	"waltuh/animal"
)

func main() {
	// Create a new Animal instance using the constructor function.
	cat := animal.NewCat("Jan")
	dog := animal.NewAnimal("Jesse")

	// Convert *animal.Cat to animal.Animal and add to the slice
	animals := []*animal.Animal{&cat.Animal, dog}

	// Access the name via the getter method, not directly.
	for i := 0; i < len(animals); i++ {
		fmt.Printf("%s\n", animals[i].GetAnimalName())
	}
}
