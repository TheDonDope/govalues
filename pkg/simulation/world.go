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
