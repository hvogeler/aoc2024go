package maze2p2

import "sort"

type ShortestPath struct {
	path []*NodeTile
	pathByPos map[Position]*NodeTile
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

	return ShortestPath{
		path: path,
		pathByPos: pm,
	}
}

func (m *Maze) WalkShortestPaths(tile *NodeTile, path []*NodeTile) {
	path = append(path, tile)
	for i, preTile := range tile.preTile {
		if i > 0 {
			newPath := path
			m.shortestPaths = append(m.shortestPaths, newPath)
            m.WalkShortestPaths(preTile, newPath)
		}
		m.WalkShortestPaths(preTile, path)
		
	}
}