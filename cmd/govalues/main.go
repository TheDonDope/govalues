package main

import (
	"fmt"
	"github.com/TheDonDope/govalues/pkg/simulation"
	"math/rand"
)

func main() {
	fmt.Println("Welcome to go values.")

	citizensNuremberg := 518365
	// Create a new world with a Nuremburg sized population
	world := &simulation.World{
		SizeX: rand.Float64(),
		SizeY: rand.Float64(),
	}
	world.Citizens = world.RandomPopulation(citizensNuremberg)

	// Start the infinite loop of life and death
	world.Run()
}
