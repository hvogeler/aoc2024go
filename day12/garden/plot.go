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

func (plot Plot) Corners() []Corner {
	corners := []Corner{}

	//   +--
	//   | x
	if plot.above().isNil() && plot.left().isNil() {
			corners = append(corners, Corner{&plot, convex, topLeft})
	}

	//   --+
	//   x |
	if plot.above().isNil() && plot.right().isNil() {
			corners = append(corners, Corner{&plot, convex, topRight})
	}

	//   x |
	//   --+
	if plot.below().isNil() && plot.right().isNil() {
			corners = append(corners, Corner{&plot, convex, bottomRight})
	}

	//   | x
	//   +--
	if plot.below().isNil() && plot.left().isNil() {
			corners = append(corners, Corner{&plot, convex, bottomLeft})
	}

	// O |
	//   +--
	// O  O
	if plot.above().isDefined() && plot.right().isDefined() && plot.above().right().isNil() && plot.right().above().isNil() {
		corners = append(corners, Corner{&plot, concave, topRight})
	}

	// O  O
	//   +--
	// O |
	if plot.below().isDefined() && plot.right().isDefined() && plot.below().right().isNil() && plot.right().below().isNil() {
		corners = append(corners, Corner{&plot, concave, bottomRight})
	}

	//    O  O
	//   --+
	//     | O
	if plot.below().isDefined() && plot.left().isDefined() && plot.below().left().isNil() && plot.left().below().isNil() {
		corners = append(corners, Corner{&plot, concave, bottomLeft})
	}

	//     | O
	//   --+
	//    O  O
	if plot.above().isDefined() && plot.left().isDefined() && plot.above().left().isNil() && plot.left().above().isNil() {
		corners = append(corners, Corner{&plot, concave, topLeft})
	}
	return corners
}

func (plot *Plot) isDefined() bool {
	return plot != nil
}

func (plot *Plot) isNil() bool {
	return plot == nil
}

func (plot Plot) above() *Plot {
	return plot.neighbors[above]
}

func (plot Plot) right() *Plot {
	return plot.neighbors[right]
}

func (plot Plot) below() *Plot {
	return plot.neighbors[below]
}

func (plot Plot) left() *Plot {
	return plot.neighbors[left]
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

// func (plot Plot) CornerType() CornerType {
// 	if plot.countNeighbors() == 1 {
// 		return deadEnd
// 	}
// 	if plot.countNeighbors() == 4 {
// 		return noCorner
// 	}
// 	if plot.countNeighbors() == 2 {
// 		if plot.neighbors[above] != nil && plot.neighbors[below] != nil {
// 			return noCorner
// 		}
// 		if plot.neighbors[left] != nil && plot.neighbors[right] != nil {
// 			return noCorner
// 		}
// 		if plot.neighbors[above] != nil && plot.neighbors[right] != nil {
// 			if plot.neighbors[above].neighbors[right] == nil {
// 				return concave
// 			} else {
// 				return convex
// 			}
// 		}
// 		if plot.neighbors[above] != nil && plot.neighbors[left] != nil {
// 			if plot.neighbors[above].neighbors[left] == nil {
// 				return concave
// 			} else {
// 				return convex
// 			}
// 		}
// 		if plot.neighbors[below] != nil && plot.neighbors[right] != nil {
// 			if plot.neighbors[below].neighbors[right] == nil {
// 				return concave
// 			} else {
// 				return convex
// 			}
// 		}
// 		if plot.neighbors[below] != nil && plot.neighbors[left] != nil {
// 			if plot.neighbors[below].neighbors[left] == nil {
// 				return concave
// 			} else {
// 				return convex
// 			}
// 		}
// 	}
// 	panic("Unexpected case of neighbor relation")

// }

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
	region.plots = append(region.plots, *plot)
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
