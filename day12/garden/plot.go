package garden

const (
	above = iota
	right
	below
	left
)

type Plot struct {
	plantType          PlantType
	location           Location
	neighbors          [4]*Plot
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
		if plot.neighbors[direction] != nil && !region.containsPlot(plot.neighbors[direction]) {
			plot.neighbors[direction].WalkPlot(region)
		}
	}
}

type PlantType rune

func (plantType PlantType) String() string {
	return string(plantType)
}
