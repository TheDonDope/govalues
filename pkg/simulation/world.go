package simulation

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"math/rand"

	log "github.com/sirupsen/logrus"

	"github.com/TheDonDope/govalues/pkg/politics"
)

// MaxReach is the maximum distance two citizens can be afar for combat interaction.
const MaxReach = 1

// Distance returns the euclidic distance between two coordinates in a two dimensional plane.
func Distance(x, y Coordinate) float64 {
	return math.Sqrt(math.Pow(x.X-y.X, 2) + math.Pow(x.Y-y.Y, 2))
}

// IsReachable returns if the given citizens are in reach for combat.
func IsReachable(c, d Citizen) bool {
	return Distance(c.Coordinate, d.Coordinate) <= MaxReach
}

// Boundary represents the upper horizontal and vertical boundaries of the world size.
type Boundary struct {
	X float64
	Y float64
}

func (b Boundary) String() string {
	return fmt.Sprintf("%.2f, %.2f", b.X, b.Y)
}

// Coordinate represents the two dimensional position within the world.
type Coordinate struct {
	X float64
	Y float64
}

func (c Coordinate) String() string {
	return fmt.Sprintf("%.2f, %.2f", c.X, c.Y)
}

func (c *Coordinate) ensureBounds(b Boundary) {
	// Check for lower x bounds
	if c.X < 0 {
		c.X = 0
	} else if c.X > b.X {
		// Check for upper x bounds
		c.X = b.X
	}
	// Check for lower y bounds
	if c.Y < 0 {
		c.Y = 0
	} else if c.Y > b.Y {
		// Check for upper y bounds
		c.Y = b.Y
	}
}

// World represents a simulation container with multiple persons.
type World struct {
	Citizens   []Citizen
	BodyBags   []Citizen
	Boundaries Boundary
}

// RandomCitizen returns a fresh all new citizen.
func (w *World) RandomCitizen() Citizen {
	c := Citizen{
		ID:        rand.Int(),
		Hitpoints: rand.Intn(MaxHitpoints),
		Coordinate: Coordinate{
			// Restrict placement of a citizen to be within the boundaries of the world.
			X: rand.Float64() * w.Boundaries.X,
			Y: rand.Float64() * w.Boundaries.Y,
		},
		Ideology: politics.Ideologies[rand.Intn(len(politics.Ideologies))],
	}

	log.WithFields(log.Fields{
		"Coord":    c.Coordinate,
		"Ideology": c.Ideology.Name,
		"HP":       c.Hitpoints,
	}).Info("Citizen created.")

	return c
}

// RandomPopulation returns a collection of citizen with the size of the given count.
func (w *World) RandomPopulation(count int) []Citizen {
	var population []Citizen

	for i := 0; i < count; i++ {
		population = append(population, w.RandomCitizen())
	}

	return population
}

// IsLastSurvivor returns true if there is only one citizen left.
func (w *World) IsLastSurvivor() bool {
	return len(w.Citizens) <= 1
}

// Shutdown logs the last surviver and dumps the state of the world at its last breath.
func (w *World) Shutdown() {
	log.WithFields(log.Fields{"Ideology": w.Citizens[0].Ideology.Name,
		"#ID": w.Citizens[0].ID}).Info("Round finished.")
	w.dump()
}

// Run is the infinite loop of life and death.
func (w *World) Run() {
	var theyLive bool = !w.IsLastSurvivor()

	for theyLive {
		// Get two random indices
		i := rand.Intn(len(w.Citizens))
		j := rand.Intn(len(w.Citizens))

		// A citizen should not shoot at him/herself
		if i == j {
			continue
		}

		// Take two random citizen
		// c := w.Citizens[i]
		// d := w.Citizens[j]

		// Let the battle commence...

		if IsReachable(w.Citizens[i], w.Citizens[j]) {
			if WillFight(w.Citizens[i], w.Citizens[j]) {
				w.Citizens[i].Conflict(w.Citizens[j], w.Boundaries)
			}
		}

		// Return or Remove oneCitizen
		if w.Citizens[i].Hitpoints == 0 {
			// update the list of killed ideloogies
			w.Citizens[j].Killed = append(w.Citizens[j].Killed, w.Citizens[i].Ideology.Name)

			// remove oneCitizen
			w.removeCitizen(i, w.Citizens[i])

			// Update anotherIndex
			if i < j {
				j--
			}

		} else {
			// oneCitizen continues to live
			w.Citizens[i].Roam(w.Boundaries)
		}

		// Return or Remove anotherCitizen
		if w.Citizens[j].Hitpoints == 0 {
			// update the list of killed ideloogies
			w.Citizens[i].Killed = append(w.Citizens[i].Killed, w.Citizens[j].Ideology.Name)

			// remove anotherCitizen
			w.removeCitizen(j, w.Citizens[j])
			// fmt.Println(fmt.Sprintf("%4v citizens left. %v(%v) died.", len(w.Citizens), anotherCitizen.Ideology.Name, anotherCitizen.ID))
		} else {
			// anotherCitizen continues to live
			w.Citizens[j].Roam(w.Boundaries)
		}
		// Check if they still live
		theyLive = !w.IsLastSurvivor()
	}

	// Leave this cruel world.
	w.Shutdown()
}

// dump data to file world_dump.json
func (w *World) dump() {
	file, _ := json.MarshalIndent(w, "", " ")
	_ = ioutil.WriteFile("world_dump.json", file, 0644)
}

// removeCitizen removes the given citizen from the given index.
func (w *World) removeCitizen(i int, c Citizen) {

	w.BodyBags = append(w.BodyBags, c)

	// https://github.com/golang/go/wiki/SliceTricks

	// Delete without preserving order
	// w.Citizens[index] = w.Citizens[len(w.Citizens)-1]
	// w.Citizens[len(w.Citizens)-1] = Citizen{}
	// w.Citizens = w.Citizens[:len(w.Citizens)-1]

	// Delete
	if i < len(w.Citizens)-1 {
		copy(w.Citizens[i:], w.Citizens[i+1:])
	}
	w.Citizens[len(w.Citizens)-1] = Citizen{}
	w.Citizens = w.Citizens[:len(w.Citizens)-1]

	log.WithFields(log.Fields{
		"#Left":  len(w.Citizens),
		"ID":     c.ID,
		"Killed": c.Ideology.Name,
	}).Info("Citizen removed.")
}
