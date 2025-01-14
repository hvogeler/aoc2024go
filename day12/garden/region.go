package garden

import (
	"math"
	"slices"
)


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

