package simulation

import (
	"math"
	"math/rand"

	"github.com/TheDonDope/govalues/pkg/politics"
)

// MaxHitpoints is the maximum hitpoints a citizen can accumulate.
const MaxHitpoints = 100

// ClosestIdeology returns the Ideology that is closest to the citizens own ideologic values.
func ClosestIdeology(c Citizen) politics.Ideology {
	var closest politics.Ideology = politics.Ideology{
		Name:       "MaxIdeology",
		Economy:    math.MaxFloat64,
		Diplomacy:  math.MaxFloat64,
		Government: math.MaxFloat64,
		Society:    math.MaxFloat64,
	}

	var dist float64 = politics.IdeologicDistance(c.Ideology, closest)
	for _, v := range politics.Ideologies {
		d := politics.IdeologicDistance(c.Ideology, v)
		if d < dist {
			dist = d
			closest = v
		}
	}

	return closest
}

// WillFight determines if the given citizens will engage in violent interaction given their political ideoligies.
func WillFight(c, d Citizen) bool {
	dist := politics.IdeologicDistance(c.Ideology, d.Ideology)
	var fight bool
	if dist < 25 {
		// A bad day can happen to anyone, so let the dice decide...
		hit := rand.Intn(100)
		if hit <= 5 {
			fight = true
		}
	} else {
		// Any distance over 50 will lead to a fight
		fight = true
	}
	return fight
}

// Citizen represents a person in a world
type Citizen struct {
	ID         int
	Hitpoints  int
	Coordinate Coordinate
	Ideology   politics.Ideology
	Killed     []string
}

// Conflict - two citiziens shooting at each other.
func (c *Citizen) Conflict(d *Citizen, b Boundary) {
	hit := rand.Intn(10)
	dmg := rand.Intn(10)

	if hit%2 == 0 {
		// All even rolls will hit anotherCitizen
		d.loseHitpoints(dmg)
		d.fleeFrom(c, b)
		c.chaseAfter(d, b)
		// c.gainHitpoints(dmg / 2)
	} else if hit%2 == 1 {
		// All uneven rolls will hit OneCitizen
		c.loseHitpoints(dmg)
		c.fleeFrom(d, b)
		d.chaseAfter(c, b)
		// d.gainHitpoints(dmg / 2)
	}
}

// Move a citizen by a given vectorCoordinate in the boundaries of the world(maximumX/Y)
func (c *Citizen) Move(v Coordinate, b Boundary) {
	// log.WithFields(log.Fields{
	// 	"Start":  c.Coordinate,
	// 	"Vector": v,
	// }).Info("Start Move")

	// Calculate the proposed new coordinate.
	d := Coordinate{
		X: c.Coordinate.X + v.X,
		Y: c.Coordinate.Y + v.Y,
	}

	// Ensure to remain in worldbounds
	d.ensureBounds(b)

	// Set new coordinates
	c.Coordinate.X = d.X
	c.Coordinate.Y = d.Y

	// log.WithFields(log.Fields{
	// 	"End": c.Coordinate,
	// }).Info("End End")
}

// Roam lets a given citizen roam around in given boundaries.
func (c *Citizen) Roam(b Boundary) {
	// Calculate a random vector.
	d := Coordinate{
		X: rand.Float64()*MaxReach - MaxReach/2,
		Y: rand.Float64()*MaxReach - MaxReach/2,
	}

	// Move the citizen within the boundaries.
	c.Move(d, b)
}

func (c *Citizen) gainHitpoints(h int) {
	// log.WithFields(log.Fields{
	// 	"Before": c.Hitpoints,
	// }).Info("Gain HP")
	v := c.Hitpoints + h
	if v > MaxHitpoints {
		c.Hitpoints = MaxHitpoints
	} else {
		c.Hitpoints = v
	}
}

func (c *Citizen) loseHitpoints(h int) {
	// log.WithFields(log.Fields{
	// 	"Before": c.Hitpoints,
	// }).Info("Lose HP")
	v := c.Hitpoints - h
	if v < 0 {
		c.Hitpoints = 0
	} else {
		c.Hitpoints = v
	}
}

func (c *Citizen) chaseAfter(d *Citizen, b Boundary) {
	// Calculate the chasing vector.
	if c.Hitpoints > 0 {
		v := Coordinate{
			X: (d.Coordinate.X - c.Coordinate.X) * rand.Float64(),
			Y: (d.Coordinate.Y - c.Coordinate.Y) * rand.Float64(),
		}
		c.Move(v, b)
	}

}

func (c *Citizen) fleeFrom(d *Citizen, b Boundary) {
	// Calculate the fleeing vector.
	if c.Hitpoints > 0 {
		v := Coordinate{
			X: -1.0 * (d.Coordinate.X - c.Coordinate.X) * rand.Float64(),
			Y: -1.0 * (d.Coordinate.Y - c.Coordinate.Y) * rand.Float64(),
		}
		c.Move(v, b)
	}
}
