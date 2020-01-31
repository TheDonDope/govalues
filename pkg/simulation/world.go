package simulation

import (
	"github.com/TheDonDope/govalues/pkg/politics"
	"math"
	"math/rand"
)

const reachable = 2

// World represents a simulation container with multiple persons.
type World struct {
	Citizens []Citizen
	SizeX    float64
	SizeY    float64
}

// Coordinate represents the two dimensional position within the world
type Coordinate struct {
	X float64 // Default 0.0
	Y float64 // Default 0.0
}

// RandomPopulation ...
func (w *World) RandomPopulation(count int) []Citizen {
	var population []Citizen
	for i := 0; i < count; i++ {
		citizen := Citizen{
			Hitpoints: rand.Intn(100),
			Coordinate: Coordinate{
				X: rand.Float64() * w.SizeX,
				Y: rand.Float64() * w.SizeY,
			},
			Ideology: politics.Ideologies[rand.Intn(len(politics.Ideologies))],
		}
		//fmt.Println(fmt.Printf("Created citizen %v", citizen))
		population = append(population, citizen)
	}
	return population
}

func distance(x, y Coordinate) float64 {
	return math.Sqrt(math.Pow(x.X-y.X, 2) + math.Pow(x.Y-y.Y, 2))
}

// Run is the infinite loop of live and death
func (w *World) Run() {
	
	// forever
	for  {
		var numberOfCitizens = len(w.Citizens)

		// take one citizen
		for i := 0; i < numberOfCitizens; i++ {
			// and anotherone
			for j := 0; j < numberOfCitizens; j++ {

				// a citizen should not shoot at him/herself
				if (i==j) {
					continue
				}

				var oneCitizen = w.Citizens[i]
				var anotherCitizen = w.Citizens[j]
				// let them fight
				if (IsReachable(oneCitizen, anotherCitizen)) {
					if (WillFight(oneCitizen, anotherCitizen)) {
						oneCitizen, anotherCitizen = Conflict(oneCitizen, anotherCitizen)
					}
				}

				w.Citizens[i] = oneCitizen
				w.Citizens[j] = anotherCitizen
			}
		}
	}
}