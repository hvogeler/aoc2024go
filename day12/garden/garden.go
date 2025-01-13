package garden

import (
	"bufio"
	"fmt"
	"math"
	"slices"
	"strings"
)

type Garden struct {
	area       [][]Plot
	regions    []Region
	dimensions Dimensions
}

func (garden Garden) FenceCost() int {
    fenceCost := 0
    for _, region := range garden.regions {
        fenceCost += region.FenceCost()
    }
    return fenceCost
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
	
	garden.dimensions = Dimensions{rowno, len(garden.area[0])}
	garden.checkNeighbors()
    garden.findRegions()
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

func (garden *Garden) findRegions() {
	for row := 0; row < garden.dimensions.rows; row++ {
		for col := 0; col < garden.dimensions.cols; col++ {
			plot := &garden.area[row][col]
			if !plot.isAssignedToRegion {
                region := new(Region)
				plot.WalkPlot(region)
                garden.regions = append(garden.regions, *region)
			}
		}
	}
}

type Region []Plot

func (region Region) PlantType() PlantType {
	return region[0].plantType
}

func (region Region) FenceCost() int {
    return len(region) * region.Perimeter()
}

func (region Region) Perimeter() int {
    perimeter := 0
    for _, plot := range region {
        for _, neighbor := range plot.neighbors {
            if neighbor == nil {
                perimeter++
            }
        }
    }
    return perimeter
}

func (region Region) contains(plot *Plot) bool {
	for _, existingPlot := range region {
		if existingPlot.Equals(*plot) {
			return true
		}
	}
	return false
}

func (region Region) containsLocation(loc Location) bool {
	for _, existingPlot := range region {
		if existingPlot.location == loc {
			return true
		}
	}
	return false
}

func (region *Region) Sort() {
	slices.SortFunc([]Plot(*region), func(a, b Plot) int {
		return a.location.Compare(b.location)
	})
}

func (region Region) String() string {
	result := ""
	tmp := region
	tmp.Sort()
	minRow := tmp[0].location.row
	maxRow := tmp[len(tmp)-1].location.row
	minCol := math.MaxInt
	maxCol := 0
	for _, plot := range tmp {
		if plot.location.col < minCol {
			minCol = plot.location.col
		}
		if plot.location.col > maxCol {
			maxCol = plot.location.col
		}
	}

	for row := minRow; row <= maxRow; row++ {
		for col := minCol; col <= maxCol; col++ {
			if region.containsLocation((Location{row, col})) {
				result += region[0].plantType.String()
			} else {
				result += "."
			}
		}
		result += "\n"
	}
	return result
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
    isAssignedToRegion bool
}

func (plot *Plot) assignToRegion() {
    plot.isAssignedToRegion = true
}

func (a Plot) Equals(b Plot) bool {
	if a.plantType == b.plantType && a.location == b.location {
		return true
	} else {
		return false
	}
}

func (plot *Plot) WalkPlot(region *Region) {
    plot.assignToRegion()
	*region = append(*region, *plot)
	for direction := above; direction <= left; direction++ {
		if plot.neighbors[direction] != nil && !region.contains(plot.neighbors[direction]) {
			plot.neighbors[direction].WalkPlot(region)
		}
	}
}

type PlantType rune

func (plantType PlantType) String() string {
	return string(plantType)
}


type Location struct {
	row int
	col int
}

func (loc Location) String() string {
	return fmt.Sprintf("(%d, %d)", loc.row, loc.col)
}

func (loca Location) Compare(locb Location) int {
	if loca == locb {
		return 0
	}
	if loca.row < locb.row {
		return -1
	}
	if loca.row > locb.row {
		return 1
	}
	if loca.col < locb.col {
		return -1
	}
	if loca.col > locb.col {
		return 1
	}
	panic("Cannot happen")
}

type Dimensions struct {
	rows int
	cols int
}
