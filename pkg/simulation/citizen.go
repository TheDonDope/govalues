package simulation

import (
	"github.com/TheDonDope/govalues/pkg/politics"
)

// Coordinate represents the two dimensional position within the world
type Coordinate struct {
	X float64 // Default 0.0
	Y float64 // Default 0.0
}

// Citizen represents a person in a world
type Citizen struct {
	Hitpoints  int
	Coordinate Coordinate
	Ideology   politics.Ideology
}
