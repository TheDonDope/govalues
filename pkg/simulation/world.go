package simulation

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"math/rand"

	"github.com/TheDonDope/govalues/pkg/politics"
)

// MaxReach is the maximum distance two citizens can be afar for combat interaction.
const MaxReach = 2

// World represents a simulation container with multiple persons.
type World struct {
	Citizens   []Citizen
	BodyBags   []Citizen
	Boundaries Boundary
}

// Coordinate represents the two dimensional position within the world.
type Coordinate struct {
	X float64
	Y float64
}

// Boundary represents the upper horizontal and vertical boundaries of the world size.
type Boundary struct {
	Coordinate
}

// Distance returns the euclidic distance between two coordinates in a two dimensional plane.
func Distance(x, y Coordinate) float64 {
	return math.Sqrt(math.Pow(x.X-y.X, 2) + math.Pow(x.Y-y.Y, 2))
}

// IsReachable returns if the given citizens are in reach for combat.
func IsReachable(oneCitizen, anotherCitizen Citizen) bool {
	overallDistance := Distance(oneCitizen.Coordinate, anotherCitizen.Coordinate)
	return overallDistance <= MaxReach
}

// dump data to file world_dump.json
func (w *World) dump() {
	file, _ := json.MarshalIndent(w, "", " ")
	_ = ioutil.WriteFile("world_dump.json", file, 0644)
}

func (w *World) removeCitizen(index int, citizen Citizen) {

	w.BodyBags = append(w.BodyBags, citizen)

	// https://github.com/golang/go/wiki/SliceTricks

	// Delete without preserving order
	// w.Citizens[index] = w.Citizens[len(w.Citizens)-1]
	// w.Citizens[len(w.Citizens)-1] = Citizen{}
	// w.Citizens = w.Citizens[:len(w.Citizens)-1]

	// Delete
	if index < len(w.Citizens)-1 {
		copy(w.Citizens[index:], w.Citizens[index+1:])
	}
	w.Citizens[len(w.Citizens)-1] = Citizen{}
	w.Citizens = w.Citizens[:len(w.Citizens)-1]
	fmt.Println(fmt.Sprintf("%4v citizens left.", len(w.Citizens)))
}

// RandomPopulation returns a collection of citizen with the size of the given count.
func (w *World) RandomPopulation(count int) []Citizen {
	var population []Citizen
	for i := 0; i < count; i++ {
		citizen := Citizen{
			ID:        i,
			Hitpoints: rand.Intn(MaxHitpoints),
			Coordinate: Coordinate{
				// Restrict placement of a citizen to be within the boundaries of the world.
				X: rand.Float64() * w.Boundaries.X,
				Y: rand.Float64() * w.Boundaries.Y,
			},
			Ideology: politics.Ideologies[rand.Intn(len(politics.Ideologies))],
		}
		population = append(population, citizen)
	}
	return population
}

// Run is the infinite loop of life and death
func (w *World) Run() {

	for {
		// Are we done yet?
		if len(w.Citizens) <= 1 {
			fmt.Println(fmt.Sprintf("%v(%v) won.", w.Citizens[0].Ideology.Name, w.Citizens[0].ID))
			w.dump()
			break
		}

		// Get two random indices
		oneIndex := rand.Intn(len(w.Citizens))
		anotherIndex := rand.Intn(len(w.Citizens))

		// A citizen should not shoot at him/herself
		if oneIndex == anotherIndex {
			continue
		}

		// Take two random citizen
		oneCitizen := w.Citizens[oneIndex]
		anotherCitizen := w.Citizens[anotherIndex]

		// Let the battle commence...
		if IsReachable(oneCitizen, anotherCitizen) {
			if WillFight(oneCitizen, anotherCitizen) {
				oneCitizen, anotherCitizen = Conflict(oneCitizen, anotherCitizen)
			}
		}

		/// Return the combatans back to the population

		// Return or Remove oneCitizen
		if oneCitizen.Hitpoints == 0 {
			// update the list of killed ideloogies
			anotherCitizen.Killed = append(anotherCitizen.Killed, oneCitizen.Ideology.Name)

			// remove oneCitizen
			w.removeCitizen(oneIndex, oneCitizen)
			// fmt.Println(fmt.Sprintf("%4v citizens left. %v(%v) died.", len(w.Citizens), oneCitizen.Ideology.Name, oneCitizen.ID))

			// Update anotherIndex
			if oneIndex < anotherIndex {
				anotherIndex = anotherIndex - 1
			}

		} else {
			// oneCitizen continues to live
			w.Citizens[oneIndex] = Roam(oneCitizen, w.Boundaries)
		}

		// Return or Remove anotherCitizen
		if anotherCitizen.Hitpoints == 0 {
			// update the list of killed ideloogies
			oneCitizen.Killed = append(oneCitizen.Killed, anotherCitizen.Ideology.Name)

			// remove anotherCitizen
			w.removeCitizen(anotherIndex, anotherCitizen)
			// fmt.Println(fmt.Sprintf("%4v citizens left. %v(%v) died.", len(w.Citizens), anotherCitizen.Ideology.Name, anotherCitizen.ID))
		} else {
			// anotherCitizen continues to live
			w.Citizens[anotherIndex] = Roam(anotherCitizen, w.Boundaries)
		}
	}
}
