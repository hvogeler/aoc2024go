package garden

import (
	"bufio"
	"fmt"
	"strings"
)

type Garden struct {
	area [][]Plot
	regionsMap map[PlantType][]Plot
	dimensions Dimensions
}

func (garden Garden) String() string {
	var s string
	for _, row := range garden.area {
		for _, plot := range row {
			s = s + plot.plantType.String()
		}
		s = s + fmt.Sprintln()
	}
	return s
}

// Creates the garden from the input string
// TODO: populate regions and Plot.fences
func GardenFromStr(data string) Garden {
	garden := new(Garden)
	scanner := bufio.NewScanner(strings.NewReader(data))
	rowno := 0
	for scanner.Scan() {
		line := scanner.Text()
		runes := []rune(line)
		row := []Plot{}
		for colno, plantType := range runes {
			plot := new(Plot)
			plot.plantType = PlantType(plantType)
			plot.location = Location{rowno, colno}
			row = append(row, *plot)
		}
		garden.area = append(garden.area, row)
		rowno++
	}
	garden.regionsMap = make(map[PlantType][]Plot)
	garden.dimensions = Dimensions{rowno, len(garden.area[0])}
	return *garden
}

type Plot struct {
	plantType PlantType
	location Location
	fences []Fence
}

func (a Plot) Equals(b Plot) bool {
	if a.plantType == b.plantType && a.location == b.location {
		return true
	} else {
		return false
	}
}

type PlantType rune
func (plantType PlantType) String() string {
	return string(plantType)
}

type Fence struct {
	fenceType FenceType
}

type FenceType int
const (
	top FenceType = iota
	bottom
	left
	right
)

type Location struct {
	row int
	col int
}

type Dimensions struct {
	rows int
	cols int
}