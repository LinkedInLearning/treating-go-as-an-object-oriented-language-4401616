package main

import (
	"fmt"
	"time"

	"github.com/fpmoles/treating-go-as-an-object-oriented-language-4401616/pets"
)

func main() {
	sleepTime := time.Now()
	//sleepTime := time.Now().Add(time.Duration(-5) * time.Hour)
	pet := pets.NewDog("Oreo", "Black and White", "Mixed", sleepTime)
	if pet.IsHungry() {
		fmt.Println(pet.Feed("steak"))
	} else {
		fmt.Println("Your animal isn't hungry, waiting")
		time.Sleep(5 * time.Second)
		fmt.Println(pet.Feed("kibble"))
	}
	fmt.Println(pet.GiveAttention("play fetch"))
}
