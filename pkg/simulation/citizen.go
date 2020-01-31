package simulation

import (
	"github.com/TheDonDope/govalues/pkg/politics"
)

// Coordinate represents the two dimensional position within the world
type Coordinate struct {
	X float64
	Y float64
}

// Citizen represents a person in a world
type Citizen struct {
	Hitpoints  int8
	Coordinate Coordinate
	Ideology   politics.Ideology
	// ContactedIdeologies []politics.Ideology     # List of ideologies which have already been in contact with
}
