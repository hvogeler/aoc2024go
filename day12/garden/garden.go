package garden

import (
	"bufio"
	"fmt"
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
				region.Sort()
                garden.regions = append(garden.regions, *region)
			}
		}
	}
}

type Dimensions struct {
	rows int
	cols int
}
