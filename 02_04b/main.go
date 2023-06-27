package main

import (
	"fmt"

	"github.com/fpmoles/treating-go-as-an-object-oriented-language-4401616/pets"
)

func main() {
	pet := pets.Dog{
		Name:  "Oreo",
		Color: "Black and White",
		Breed: "Mixed",
	}
	fmt.Println(pet.Feed("steak"))
	fmt.Println(pet.GiveAttention("play fetch"))
}
