package gallons

import (
	"fmt"
	"math"
)

// Group struct that stores bottles and the watter
type Group struct {
	Watter  uint32
	Bottles []Bottle
}

// NewGroup return a new Group struct
func NewGroup(bootles []Bottle, watter uint32) Group {
	return Group{
		Watter:  watter,
		Bottles: bootles,
	}
}

// CreateBottleGroup based on parameters
func CreateBottleGroup(watterVolume uint32, volumes ...uint32) Group {
	bottles := make([]Bottle, len(volumes))

	for i, v := range volumes {
		bottles[i] = NewBottle(v)
	}

	return NewGroup(bottles, watterVolume)
}

// UseBottles use all the available bottles using combination to search for the one with least rest of watter
func (g Group) UseBottles() (int, []Bottle, error) {
	mapBottles := make(map[int][]Bottle)

	// generate a copy of bottles
	bottles := make([]Bottle, len(g.Bottles))
	copy(bottles, g.Bottles)

	g.addBottle(mapBottles, []Bottle{}, bottles, len(g.Bottles))

	watterRest := math.MaxInt64
	for rest := range mapBottles {
		if rest >= 0 && rest < watterRest {
			watterRest = rest
		}
	}

	bottles, ok := mapBottles[watterRest]
	if !ok {
		return 0, []Bottle{}, fmt.Errorf("Error on get the bottles")
	}

	return watterRest, bottles, nil
}

func (g Group) addBottle(mapBottles map[int][]Bottle, newCombo []Bottle, unusedBottles []Bottle, length int) {
	if length <= 0 {
		return
	}

	for _, bottle := range unusedBottles {
		if !g.valid(bottle, newCombo) {
			return
		}

		newCombo := append(newCombo, bottle)
		volume := volume(newCombo)

		rest := volume - int(g.Watter)
		_, ok := mapBottles[rest]
		if !ok {
			mapBottles[rest] = newCombo
		}

		if rest == 0 {
			return
		}

		// remove the bottle from the slice of unused bottles
		unusedBottles = removeBottle(bottle, unusedBottles)
		g.addBottle(mapBottles, newCombo, unusedBottles, length-1)

		// add the bottle again to try all the possibilities
		unusedBottles = append(unusedBottles, bottle)
	}
}

// mapBottles generates a map of the bottles map[volume]qnt
func mapBottles(bottles []Bottle) map[uint32]int {
	bottleMap := make(map[uint32]int, len(bottles))
	for _, bottle := range bottles {
		_, ok := bottleMap[bottle.VolumeML]
		if !ok {
			bottleMap[bottle.VolumeML] = 1
			continue
		}
		bottleMap[bottle.VolumeML]++
	}
	return bottleMap
}

// valid verify if it's valid to continue
func (g Group) valid(bottle Bottle, combo []Bottle) bool {
	myMap := mapBottles(g.Bottles)

	v, ok := myMap[bottle.VolumeML]
	if !ok || v <= 0 {
		return false
	}

	myMap[bottle.VolumeML]--
	for _, bottle := range combo {
		v, ok := myMap[bottle.VolumeML]
		if !ok || v <= 0 {
			return false
		}
		myMap[bottle.VolumeML]--
	}

	return true
}

// removeBottle remove the bottle from the slice
func removeBottle(bottleToRemove Bottle, bootles []Bottle) []Bottle {
	for i, bottle := range bootles {
		if bottle.VolumeML == bottleToRemove.VolumeML {
			bootles[i] = bootles[len(bootles)-1]
			bootles = bootles[:len(bootles)-1]
			return bootles
		}
	}

	return bootles
}

// volume returns the volume sum from bottles
func volume(bottles []Bottle) int {
	result := 0

	for _, b := range bottles {
		result += int(b.VolumeML)
	}

	return result
}
