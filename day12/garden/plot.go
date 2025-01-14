package garden

const (
	above = iota
	right
	below
	left
)

type CornerType int

const (
	noCorner CornerType = iota
	convex
	concave
	both
	deadEnd
)

func (cornerType CornerType) Weight() int {
	switch cornerType {
	case convex, concave:
		return 1
	case both, deadEnd:
		return 2
	}
	panic("Switch exhausted")
}

type Plot struct {
	plantType          PlantType
	location           Location
	neighbors          [4]*Plot
	cornerType         CornerType
	isAssignedToRegion bool
}

func (plot *Plot) assignToRegion() {
	plot.isAssignedToRegion = true
}

// func (plot *Plot) assignCornerType() {
// 	if plot.countNeighbors() == 1 {
// 		plot.cornerType = deadEnd
// 	}
// 	if plot.countNeighbors() == 4 {
// 		plot.cornerType = noCorner
// 	}
// 	if plot.countNeighbors() == 2 {
// 		if plot.TwoNeighborRelation() == opposite {
// 			plot.cornerType = noCorner
// 		}
// 		if plot.TwoNeighborRelation() == adjacent {
// 			plot.cornerType = convex
// 		}
// 	}
// }

// type neighborRelation int

// const (
// 	adjacent neighborRelation = iota
// 	opposite
// )

func (plot Plot) CornerType() CornerType {
	if plot.countNeighbors() == 1 {
		return deadEnd
	}
	if plot.countNeighbors() == 4 {
		return noCorner
	}
	if plot.countNeighbors() == 2 {
		if plot.neighbors[above] != nil && plot.neighbors[below] != nil {
			return noCorner
		}
		if plot.neighbors[left] != nil && plot.neighbors[right] != nil {
			return noCorner
		}
		if plot.neighbors[above] != nil && plot.neighbors[right] != nil {
			if plot.neighbors[above].neighbors[right] == nil {
				return concave
			} else {
				return convex
			}
		}
		if plot.neighbors[above] != nil && plot.neighbors[left] != nil {
			if plot.neighbors[above].neighbors[left] == nil {
				return concave
			} else {
				return convex
			}
		}
		if plot.neighbors[below] != nil && plot.neighbors[right] != nil {
			if plot.neighbors[below].neighbors[right] == nil {
				return concave
			} else {
				return convex
			}
		}
		if plot.neighbors[below] != nil && plot.neighbors[left] != nil {
			if plot.neighbors[below].neighbors[left] == nil {
				return concave
			} else {
				return convex
			}
		}
	}
	panic("Unexpected case of neighbor relation")

}

func absInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func (plot Plot) countNeighbors() int {
	sum := 0
	for _, neighbor := range plot.neighbors {
		if neighbor != nil {
			sum++
		}
	}
	return sum
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
