package simulation

import (
	"github.com/TheDonDope/govalues/pkg/politics"
)

// Citizen represents a person in a world
type Citizen struct {
	Hitpoints int8
	Ideology  politics.Ideology
}
