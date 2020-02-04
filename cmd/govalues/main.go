package main

import (
	"fmt"
	"github.com/TheDonDope/govalues/pkg/simulation"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to go values.")
	rand.Seed(time.Now().UTC().UnixNano())

	// citizensNuremberg := 518365
	// citizensNurembergInnerCity := 51836
	// Create a new world with a Nuremburg sized population
	world := &simulation.World{
		Boundaries: simulation.Coordinate{
			// Restrict placement of a citizen to be within the boundaries of the world.
			X: 100,
			Y: 100,
		},
	}
	world.Citizens = world.RandomPopulation(100)

	// Start the infinite loop of life and death
	world.Run()
}
