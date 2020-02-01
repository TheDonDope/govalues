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
