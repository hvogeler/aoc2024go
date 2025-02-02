package maze2

import "sort"

type ShortestPath struct {
	path []*NodeTile
	pathByPos map[Position]*NodeTile
}

func NewShortestPath(m *Maze) ShortestPath {
	path := []*NodeTile{}
	tile := m.finishTile
	for tile.preTile != nil {
		path = append(path, tile)
		tile = tile.preTile
	}
	path = append(path, tile)

	sort.SliceStable(path, func(i, j int) bool {
		return j < i
	})

	pm := make(map[Position]*NodeTile)
	for _, tile := range path {
		pm[tile.pos] = tile
	}

	return ShortestPath{
		path: path,
		pathByPos: pm,
	}
}