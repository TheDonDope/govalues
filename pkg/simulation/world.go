package simulation

import (
	"github.com/TheDonDope/govalues/pkg/politics"
	"math"
	"math/rand"
)

// MaxReach is the maximum distance two citizens can be afar for combat interaction.
const MaxReach = 2

// World represents a simulation container with multiple persons.
type World struct {
	Citizens   []Citizen
	Boundaries Coordinate
}

// Coordinate represents the two dimensional position within the world.
type Coordinate struct {
	X float64
	Y float64
}

// Distance returns the euclidic distance between two coordinates in a two dimensional plane.
func Distance(x, y Coordinate) float64 {
	return math.Sqrt(math.Pow(x.X-y.X, 2) + math.Pow(x.Y-y.Y, 2))
}

// IsReachable returns if the given citizens are in reach for combat.
func IsReachable(oneCitizen, anotherCitizen Citizen) bool {
	overallDistance := Distance(oneCitizen.Coordinate, anotherCitizen.Coordinate)
	return overallDistance <= MaxReach
}

// RandomPopulation returns a collection of citizen with the size of the given count.
func (w *World) RandomPopulation(count int) []Citizen {
	var population []Citizen
	for i := 0; i < count; i++ {
		citizen := Citizen{
			Hitpoints: rand.Intn(100),
			Coordinate: Coordinate{
				// Restrict placement of a citizen to be within the boundaries of the world.
				X: rand.Float64() * w.Boundaries.X,
				Y: rand.Float64() * w.Boundaries.Y,
			},
			Ideology: politics.Ideologies[rand.Intn(len(politics.Ideologies))],
		}
		population = append(population, citizen)
	}
	return population
}

// Run is the infinite loop of live and death
func (w *World) Run() {

	for {
		// Take two random citizen
		oneIndex := rand.Intn(len(w.Citizens))
		anotherIndex := rand.Intn(len(w.Citizens))
		oneCitizen := w.Citizens[oneIndex]
		anotherCitizen := w.Citizens[anotherIndex]

		// Do not beat on a dead horse
		if anotherCitizen.Hitpoints <= 0 || oneCitizen.Hitpoints <= 0 {
			// TODO: Remove dead horses from the fight equation
			continue
		}

		// A citizen should not shoot at him/herself
		if oneCitizen == anotherCitizen {
			continue
		}

		// Let the battle commence...
		if IsReachable(oneCitizen, anotherCitizen) {
			if WillFight(oneCitizen, anotherCitizen) {
				oneCitizen, anotherCitizen = Conflict(oneCitizen, anotherCitizen)
			}
		}
		// Return the combatans back to the population
		w.Citizens[oneIndex] = Roam(oneCitizen, w.Boundaries)
		w.Citizens[anotherIndex] = Roam(anotherCitizen, w.Boundaries)
	}

}
