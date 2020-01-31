package simulation

import (
	"fmt"
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

// IsReachable returns if the given citizens are in reach for combat.
func IsReachable(oneCitizen, anotherCitizen Citizen) bool {
	overallDistance := math.Sqrt(math.Pow(anotherCitizen.Coordinate.X-oneCitizen.Coordinate.X, 2) + math.Pow(anotherCitizen.Coordinate.Y-oneCitizen.Coordinate.Y, 2))
	return overallDistance <= reachable
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
		fmt.Println(fmt.Printf("Created citizen %v", citizen))
		population = append(population, citizen)
	}
	return population
}
