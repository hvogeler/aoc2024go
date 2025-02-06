package maze2p2

import (
	"bufio"
	"container/heap"
	"fmt"
	"sort"
	"strings"
)

type Maze struct {
	tiles         map[Position]Tile
	startTile     *NodeTile
	finishTile    *NodeTile
	dimensions    Dimensions
	shortestPaths [][]*NodeTile
}

// findPath will set the cost of the Finish Tile to the cost of the shortset path.
// It will also point back from the Finish Tile to the preceeding tile.

// https://www.reddit.com/r/adventofcode/comments/1hfmbel/2024_day_16_part_1_c_stuck_in_part_1/
// You need to allow entering some cell facing a given direction, even if you've already reached
// the same cell facing a different direction. So, the nodes in the graph you're searching should
// include the direction you're facing. I.e. your distance array should be distance[x][y][dir]
func (m *Maze) FindPath() {
	srcNode := m.startTile
	pq := &NodeHeap{}
	heap.Init(pq)
	for {
		// Process Phase
		for _, heading := range Headings() {
			if srcNode.heading.IsOpposite(heading) {
                // Don't go backwards. Its is a DAG
				continue
			}
			neighbor := m.NeighborTile(*srcNode, heading)
			if neighbor.TileType() == WallType {
				continue
			}
			if neighbor.TileType() == NodeType {
				tgtNode := neighbor.(*NodeTile)
				cost := srcNode.cost + srcNode.heading.Cost(heading) + 1
				tgtCost := tgtNode.cost
				if cost < tgtCost {
                    // Relax node
					tgtNode.cost = cost
					tgtNode.preTile = []*NodeTile{srcNode}
					tgtNode.heading = heading
					heap.Push(pq, tgtNode)
					if tgtNode.pos.row == m.finishTile.pos.row && tgtNode.pos.col == m.finishTile.pos.col {
						tgtNode.nodeType = Finish
						if tgtNode.cost < m.finishTile.cost {
							m.finishTile = tgtNode
						}
					}
				}
				if cost == tgtCost {
                    // Add new possible shortest path to node
					tgtNode.preTile = append(tgtNode.preTile, srcNode)
					tgtNode.heading = heading
					heap.Push(pq, tgtNode)
				}
				// tgtNode.Println(true)
			}
		}
		srcNode.isExplored = true
		// Identification Phase
		if pq.Len() == 0 {
			break
		}
		srcNode = heap.Pop(pq).(*NodeTile)
		for pq.Len() > 0 && srcNode.isExplored {
			srcNode = heap.Pop(pq).(*NodeTile)
		}
	}
}

func (m Maze) Score() int {
	return m.finishTile.cost
}

func (m *Maze) NeighborTile(tile NodeTile, h Heading) Tile {
	var nPos Position
	switch h {
	case North:
		nPos = NewPosition(tile.pos.row-1, tile.pos.col, h)
	case East:
		nPos = NewPosition(tile.pos.row, tile.pos.col+1, h)
	case South:
		nPos = NewPosition(tile.pos.row+1, tile.pos.col, h)
	case West:
		nPos = NewPosition(tile.pos.row, tile.pos.col-1, h)
	default:
		panic("Switch exhausted")
	}
	ntile, exists := m.Tile(nPos.row, nPos.col, nPos.heading)
	if exists {
		return ntile
	} else {
		// create new node if getting to a node from a different heading.
		// Nodes have different cost depending whether they were approached straight ahead
		// or from the side
		// case: >o> -  o cost is 1
		// case: o>
		//       ^   -  o costs 1000
		m.tiles[NewPosition(nPos.row, nPos.col, nPos.heading)] = NewNodeTile(nPos.row, nPos.col, nPos.heading)
		return m.tiles[NewPosition(nPos.row, nPos.col, nPos.heading)]
	}
}

func (m Maze) Tile(row int, col int, h Heading) (Tile, bool) {
	pos := NewPosition(row, col, h)
	tile, exists := m.tiles[pos]
	if exists {
		return tile, true
	} else {
		if row < 0 || row >= m.dimensions.rows || col < 0 || col >= m.dimensions.cols {
			panic("Off Maze Grid")
		}
		if walltile, wallexists := m.tiles[NewPosition(row, col, Undefined)]; wallexists {
			wt, ok := walltile.(WallTile)
			if ok {
				return wt, true
			}
		}
		return nil, false
	}
}

func (m Maze) AnyTile(row int, col int, sp ShortestPath) Tile {
	if pathTile, ok := sp.pathByLoc[NewLocation(row, col)]; ok {
		return pathTile
	}
	if wall, ok := m.tiles[NewPosition(row, col, Undefined)]; ok {
		return wall
	}
	return NewNodeTile(row, col, Undefined)
}

func (m Maze) String() string {
	var s strings.Builder
	for row := 0; row < m.dimensions.rows; row++ {
		s.WriteString(fmt.Sprintf("%4d. ", row))
		for col := 0; col < m.dimensions.cols; col++ {
			if tile, exists := m.Tile(row, col, Undefined); exists {
            			s.WriteString(fmt.Sprint(tile))
            } else {
                s.WriteString(fmt.Sprint(Undefined))
            }
		}
		s.WriteString(fmt.Sprintln())
	}
	s.WriteString(fmt.Sprintf("Number of Tiles: %d\n", len(m.tiles)))
	return s.String()
}

func (m Maze) ShortestPaths() [][]*NodeTile {
	return m.shortestPaths
}

func (m Maze) FinishTile() *NodeTile {
	return m.finishTile
}

func (m Maze) CountAllVisitedTiles() int {
	tileMap := make(map[Location]bool)
	for _, path := range m.ShortestPaths() {
		for _, tile := range path {
			tileMap[NewLocation(tile.pos.row, tile.pos.col)] = true
		}
	}
	tiles := []Location{}
	for k := range tileMap {
		tiles = append(tiles, k)
	}

	sort.SliceStable(tiles, func(i, j int) bool {
		if tiles[i].row < tiles[j].row {
			return true
		}
		if tiles[i].row == tiles[j].row {
			if tiles[i].col < tiles[j].col {
				return true
			}
		}
		return false
	})

	return len(tiles)
}

func (m *Maze) WalkShortestPaths(finishTile *NodeTile, path []*NodeTile) {
	path = append(path, finishTile)
	if finishTile.nodeType == Start {
		p := make([]*NodeTile, len(path))
		copy(p, path)
		m.shortestPaths = append(m.shortestPaths, p)
		return
	}
	for i, preTile := range finishTile.preTile {
		if i > 0 {
			newPath := path
			m.WalkShortestPaths(preTile, newPath)
		} else {
			m.WalkShortestPaths(preTile, path)
		}
	}
}

func (m Maze) PrintPath(path []*NodeTile) string {
	var s strings.Builder
	sp := NewShortestPath(path)
	for row := 0; row < m.dimensions.rows; row++ {
		s.WriteString(fmt.Sprintf("%4d. ", row))
		for col := 0; col < m.dimensions.cols; col++ {
			if tile, ok := sp.pathByLoc[NewLocation(row, col)]; ok {
				s.WriteString(fmt.Sprint(tile))
				continue
			}
			if tile, ok := m.tiles[NewPosition(row, col, Undefined)]; ok {
				s.WriteString(fmt.Sprint(tile))
				continue
			}
			s.WriteString(fmt.Sprint(Undefined))
		}
		s.WriteString(fmt.Sprintln())
	}
	s.WriteString(fmt.Sprintf("Number of Tiles: %d\n", len(m.tiles)))
	return s.String()
}

func MazeFromStr(s string) Maze {
	newMaze := Maze{}
	tiles := make(map[Position]Tile)
	scanner := bufio.NewScanner(strings.NewReader(s))
	var cols int
	var rowno int
	for rowno = 0; scanner.Scan(); rowno++ {
		line := scanner.Text()
		if cols == 0 {
			cols = len(line)
		}
		for colno, tileChar := range strings.Split(line, "") {
			switch TileType(tileChar) {
			case WallType:
				tile := WallTile{}
				tiles[NewPosition(rowno, colno, Undefined)] = tile
			case StartType:
				newNode := NewNodeTile(rowno, colno, East)
				newNode.nodeType = Start
				newNode.cost = 0
				newMaze.startTile = newNode
			case FinishType:
				newNode := NewNodeTile(rowno, colno, Undefined)
				newNode.nodeType = Finish
				newMaze.finishTile = newNode
			}
		}
	}
	newMaze.tiles = tiles
	newMaze.dimensions = Dimensions{
		rows: rowno,
		cols: cols}
	return newMaze
}

type Dimensions struct {
	rows int
	cols int
}
