package simulation

import (
	"fmt"
	"github.com/TheDonDope/govalues/pkg/politics"
	"math"
	"math/rand"
)

// Citizen represents a person in a world
type Citizen struct {
	Hitpoints  int
	Coordinate Coordinate
	Ideology   politics.Ideology
}

// ClosestIdeology returns the Ideology that is closest to the citizens own ideologic values.
func ClosestIdeology(c Citizen) politics.Ideology {
	var closestIdeology politics.Ideology = politics.Ideology{
		Name:       "MaxIdeology",
		Economy:    math.MaxFloat64,
		Diplomacy:  math.MaxFloat64,
		Government: math.MaxFloat64,
		Society:    math.MaxFloat64,
	}
	var closestDistance float64 = politics.IdeologicDistance(c.Ideology, closestIdeology)

	for _, v := range politics.Ideologies {
		d := politics.IdeologicDistance(c.Ideology, v)
		if d < closestDistance {
			closestDistance = d
			closestIdeology = v
		}
	}

	return closestIdeology
}

// WillFight determines if the given citizens will engage in violent interaction given their political ideoligies.
func WillFight(oneCitizen, anotherCitizen Citizen) bool {
	distance := politics.IdeologicDistance(oneCitizen.Ideology, anotherCitizen.Ideology)
	var fightEnsured bool
	if distance < 50 {
		// A bad day can happen to anyone, so let the dice decide...
		d100Roll := rand.Intn(100)
		if d100Roll <= 5 {
			fightEnsured = true
		}
	} else {
		// Any distance over 50 will lead to a fight
		fightEnsured = true
	}
	return fightEnsured
}

// Conflict - two citiziens shooting at each other. Returns both citizens after the fighting is done
func Conflict(oneCitizen, anotherCitizen Citizen) (Citizen, Citizen) {
	fmt.Println(oneCitizen.Ideology.Name + " vs. " + anotherCitizen.Ideology.Name)
	hitRoll := rand.Intn(10)
	dmgRoll := rand.Intn(100)

	if hitRoll%2 == 0 {
		// All even rolls will hit anotherCitizen
		anotherCitizen.Hitpoints -= dmgRoll
	} else if hitRoll%2 == 1 {
		// All uneven rolls will hit OneCitizen
		oneCitizen.Hitpoints -= dmgRoll
	}

	return oneCitizen, anotherCitizen
}

// Move a citizen by a given vectorCoordinate in the boundaries of the world(maximumX/Y)
func Move(citizen Citizen, vectorCoordinate Coordinate, boundaries Coordinate) Citizen {
	
	// calculate the proposed new coordinate
	newCoordinate := Coordinate{
		X: citizen.Coordinate.X + vectorCoordinate.X,
		Y: citizen.Coordinate.Y + vectorCoordinate.Y,
	}

	// check whether the proposed new coordinate is still in the boundaries 
	if newCoordinate.X < boundaries.X {
		citizen.Coordinate.X = newCoordinate.X
	} else {
		citizen.Coordinate.X = boundaries.X
	}
	if newCoordinate.Y < boundaries.Y {
		citizen.Coordinate.Y = newCoordinate.Y
	} else {
		citizen.Coordinate.Y = boundaries.Y
	}

	return citizen
}

// Roam lets a given citizen roam around in given boundaries
func Roam(citizen Citizen, boundaries Coordinate) Citizen {

	// calculate a random vector
	vectorCoordinate := Coordinate{
		X: rand.Float64() * MaxReach,
		Y: rand.Float64() * MaxReach,
	}

	return Move(citizen, vectorCoordinate, boundaries)
}
