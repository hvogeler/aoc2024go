package garden

import (
	"math"
	"slices"
)

type Region struct {
	plots []Plot
	// corners []Corner
}

func (region Region) FenceCost() int {
	return region.Length() * region.Perimeter()
}

func (region Region) Perimeter() int {
	perimeter := 0
	for _, plot := range region.plots {
		for _, neighbor := range plot.neighbors {
			if neighbor == nil {
				perimeter++
			}
		}
	}
	return perimeter
}

func (region Region) Sides() int {
	return 42
}

func (region Region) Length() int {
	return len(region.plots)
}

func (region Region) PlantType() PlantType {
	return region.plots[0].plantType
}

func (region Region) containsPlot(plot *Plot) bool {
	for _, existingPlot := range region.plots {
		if existingPlot.Equals(*plot) {
			return true
		}
	}
	return false
}

func (region Region) containsLocation(loc Location) bool {
	for _, existingPlot := range region.plots {
		if existingPlot.location == loc {
			return true
		}
	}
	return false
}

func (region *Region) Sort() {
	slices.SortFunc(region.plots, func(a, b Plot) int {
		return a.location.Compare(b.location)
	})
}

func (region Region) String() string {
	result := ""
	minRow := region.plots[0].location.row
	maxRow := region.plots[region.Length()-1].location.row
	minCol := math.MaxInt
	maxCol := 0
	for _, plot := range region.plots {
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
				result += region.plots[0].plantType.String()
			} else {
				result += "."
			}
		}
		result += "\n"
	}
	return result
}
