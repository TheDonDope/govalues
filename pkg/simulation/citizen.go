package simulation

import (
	"fmt"
	"github.com/TheDonDope/govalues/pkg/politics"
	"math/rand"
)

// Citizen represents a person in a world
type Citizen struct {
	Hitpoints  int
	Coordinate Coordinate
	Ideology   politics.Ideology
}

// IsReachable returns if the given citizens are in reach for combat.
func IsReachable(oneCitizen, anotherCitizen Citizen) bool {
	overallDistance := distance(oneCitizen.Coordinate, anotherCitizen.Coordinate)
	return overallDistance <= reachable
}

// WillFight determines if the given citizens will engage in violent interaction given their political ideoligies.
func WillFight(oneCitizen, anotherCitizen Citizen) bool {
	distance := politics.IdeologicDistance(oneCitizen.Ideology, anotherCitizen.Ideology)
	var fightEnsured bool
	if distance >= 50 {
		fightEnsured = true
	} else if distance < 50 {
		// Jeder hat mal einen schlechten Tag -> Die Wuerfel entscheiden...
		d100Roll := rand.Intn(100)
		if d100Roll <= 5 {
			fightEnsured = true
		}
	}
	return fightEnsured
}

// Conflict - two citiziens shooting at each other
// returns true if the first citizien this the otherCitizien
func Conflict(oneCitizen, anotherCitizen Citizen) (Citizen, Citizen, ) {

	fmt.Println(oneCitizen.Ideology.Name + " vs. "+ anotherCitizen.Ideology.Name)
	return oneCitizen, anotherCitizen
}