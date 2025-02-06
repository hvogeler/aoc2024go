package maze2p2

import "sort"

type Location struct {
	row int
	col int
}

func NewLocation(row int, col int) Location {
	return Location{row, col}
}

type ShortestPath struct {
	path []*NodeTile
	pathByPos map[Position]*NodeTile
	pathByLoc map[Location]*NodeTile
}

func NewShortestPath(path []*NodeTile) ShortestPath {
	// path := []*NodeTile{}
	// tile := m.finishTile
	// for tile.preTile != nil {
	// 	path = append(path, tile)
	// 	//TODO: Walk all paths here
	// 	tile = tile.preTile[0]
	// }
	// path = append(path, tile)

	sort.SliceStable(path, func(i, j int) bool {
		return j < i
	})

	pm := make(map[Position]*NodeTile)
	for _, tile := range path {
		pm[tile.pos] = tile
	}

	lm := make(map[Location]*NodeTile)
	for _, tile := range path {
		lm[NewLocation(tile.pos.row, tile.pos.col)] = tile
	}

	return ShortestPath{
		path: path,
		pathByPos: pm,
		pathByLoc: lm,
	}
}
