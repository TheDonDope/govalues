package simulation

import (
	"math"
)

const reachable = 2

// World represents a simulation container with multiple persons.
type World struct {
	Citizens []Citizen
}

// IsReachable returns if the given citizens are in reach for combat.
func IsReachable(oneCitizen, anotherCitizen Citizen) bool {
	overallDistance := math.Sqrt(math.Pow(anotherCitizen.Coordinate.X-oneCitizen.Coordinate.X, 2) + math.Pow(anotherCitizen.Coordinate.Y-oneCitizen.Coordinate.Y, 2))
	return overallDistance <= reachable
}
