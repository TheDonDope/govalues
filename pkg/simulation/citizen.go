package simulation

import (
	"math"
	"math/rand"

	"github.com/TheDonDope/govalues/pkg/politics"
)

// MaxHitpoints is the maximum hitpoints a citizen can accumulate.
const MaxHitpoints = 100

// Citizen represents a person in a world
type Citizen struct {
	ID         int
	Hitpoints  int
	Coordinate Coordinate
	Ideology   politics.Ideology
	Killed     []string
}

func (c *Citizen) gainHitpoints(hitpoints int) {
	newHitpoints := c.Hitpoints + hitpoints
	if newHitpoints > MaxHitpoints {
		c.Hitpoints = MaxHitpoints
	} else {
		c.Hitpoints = newHitpoints
	}
}

func (c *Citizen) loseHitpoints(hitpoints int) {
	newHitpoints := c.Hitpoints - hitpoints
	if newHitpoints < 0 {
		c.Hitpoints = 0
	} else {
		c.Hitpoints = newHitpoints
	}
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
	if distance < 25 {
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
	// fmt.Println(fmt.Sprintf("%v(%v) vs. %v(%v) ", oneCitizen.Ideology.Name, oneCitizen.ID, anotherCitizen.Ideology.Name, anotherCitizen.ID))
	hitRoll := rand.Intn(10)
	dmgRoll := rand.Intn(10)

	if hitRoll%2 == 0 {
		// All even rolls will hit anotherCitizen
		anotherCitizen.loseHitpoints(dmgRoll)
		oneCitizen.gainHitpoints(dmgRoll / 2)
	} else if hitRoll%2 == 1 {
		// All uneven rolls will hit OneCitizen
		oneCitizen.loseHitpoints(dmgRoll)
		anotherCitizen.gainHitpoints(dmgRoll / 2)
	}

	return oneCitizen, anotherCitizen
}

// Move a citizen by a given vectorCoordinate in the boundaries of the world(maximumX/Y)
func Move(citizen Citizen, vectorCoordinate Coordinate, boundaries Boundary) Citizen {

	// calculate the proposed new coordinate
	newCoordinate := Coordinate{
		X: citizen.Coordinate.X + vectorCoordinate.X,
		Y: citizen.Coordinate.Y + vectorCoordinate.Y,
	}

	// check whether the proposed new coordinate is still in the boundaries
	if newCoordinate.X < 0.0 {
		newCoordinate.X = 0.0
	} else if newCoordinate.X > boundaries.X {
		citizen.Coordinate.X = boundaries.X
	} else {
		citizen.Coordinate.X = newCoordinate.X
	}

	if newCoordinate.Y < 0.0 {
		newCoordinate.Y = 0.0
	} else if newCoordinate.Y > boundaries.Y {
		citizen.Coordinate.Y = boundaries.Y
	} else {
		citizen.Coordinate.Y = newCoordinate.Y
	}

	return citizen
}

// Roam lets a given citizen roam around in given boundaries
func Roam(citizen Citizen, boundaries Boundary) Citizen {

	// calculate a random vector
	vectorCoordinate := Coordinate{
		X: rand.Float64()*MaxReach - MaxReach/2,
		Y: rand.Float64()*MaxReach - MaxReach/2,
	}

	return Move(citizen, vectorCoordinate, boundaries)
}
