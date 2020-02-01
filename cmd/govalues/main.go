package main

import (
	"fmt"
	"github.com/TheDonDope/govalues/pkg/simulation"
	"math/rand"
)

func main() {
	fmt.Println("Welcome to go values.")

	// Create a new world with random? or fixed number of citizens

	world := &simulation.World{
		SizeX: rand.Float64(),
		SizeY: rand.Float64(),
	}
	world.Citizens = world.RandomPopulation(5000)

	// start the infinite loop
	world.Run()
}
