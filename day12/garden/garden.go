package garden

import (
	"bufio"
	"fmt"
	"strings"
)

type Garden struct {
	area       [][]Plot
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
	garden.checkNeighbors()
	return *garden
}

func (garden *Garden) checkNeighbors() {
	for row := 0; row < garden.dimensions.rows; row++ {
		for col := 0; col < garden.dimensions.cols; col++ {
			plot := &garden.area[row][col]

			plotRow := plot.location.row
			plotCol := plot.location.col
			// check above
			if plotRow > 0 && garden.area[plotRow-1][plotCol].plantType == plot.plantType {
				plot.neighbors[above] = &garden.area[plotRow-1][plotCol]
			}
			// check below
			if plotRow < garden.dimensions.rows-1 && garden.area[plotRow+1][plotCol].plantType == plot.plantType {
				plot.neighbors[below] = &garden.area[plotRow+1][plotCol]
			}
			// check right
			if plotCol < garden.dimensions.cols-1 && garden.area[plotRow][plotCol+1].plantType == plot.plantType {
				plot.neighbors[right] = &garden.area[plotRow][plotCol+1]
			}
			// check left
			if plotCol > 0 && garden.area[plotRow][plotCol-1].plantType == plot.plantType {
				plot.neighbors[left] = &garden.area[plotRow][plotCol-1]
			}
		}
	}
}

// func (garden *Garden) findRegions() {
// 	for row := 0; row < garden.dimensions.rows; row++ {
// 		for col := 0; col < garden.dimensions.cols; col++ {
// 			plot := garden.area[row][col]
// 			if _, regionExists := garden.regionsMap[plot.plantType]; !regionExists {
// 				garden.exploreRegion(plot)
// 			}
// 		}
// 	}
// }

// func (garden Garden) exploreRegion(plot Plot) {

// }

type Region []Plot

func (region Region) contains(plot *Plot) bool {
	for _, existingPlot := range region {
		if existingPlot.Equals(*plot) {
			return true
		}
	}
	return false
}

func (plot Plot) WalkPlot(region *Region) {
	*region = append(*region, plot)
	for direction := above; direction <= left; direction++ {
		if plot.neighbors[direction] != nil && !region.contains(plot.neighbors[direction]) {
			plot.neighbors[direction].WalkPlot(region)
		}
	}
}

const (
	above = iota
	right
	below
	left
)

type Plot struct {
	plantType PlantType
	location  Location
	neighbors [4]*Plot
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

// type Fence struct {
// 	fenceType FenceType
// }

// type FenceType int
// const (
// 	top FenceType = iota
// 	bottom
// 	left
// 	right
// )

type Location struct {
	row int
	col int
}

func (loc Location) String() string {
	return fmt.Sprintf("(%d, %d)", loc.row, loc.col)
}

type Dimensions struct {
	rows int
	cols int
}
