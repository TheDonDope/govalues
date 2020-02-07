package main

import (
	"math/rand"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/TheDonDope/govalues/pkg/simulation"
)

func main() {
	// Roll the dice...
	rand.Seed(time.Now().UTC().UnixNano())

	//citizensNuremberg := 5183
	// citizensNurembergInnerCity := 51836
	// Create a new world with a Nuremburg sized population
	world := &simulation.World{
		Boundaries: simulation.Boundary{
			// Restrict placement of a citizen to be within the boundaries of the world.
			X: 100,
			Y: 100,
		},
	}
	size := 118
	world.Citizens = world.RandomPopulation(size)

	log.WithFields(log.Fields{"Boundaries (X Y)": world.Boundaries,
		"#Citizens": len(world.Citizens)}).Info("World created.")

	// Start the infinite loop of life and death
	world.Run()
	// world.Shutdown()
}
