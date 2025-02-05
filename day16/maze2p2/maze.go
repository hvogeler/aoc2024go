package maze2p2

import (
	"bufio"
	"container/heap"
	"fmt"
	"strings"
)

type Maze struct {
	tiles      map[Position]Tile
	startTile  *NodeTile
	finishTile *NodeTile
	dimensions Dimensions
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
			neighbor := m.NeighborTile(*srcNode, heading)
			if neighbor.TileType() == WallType || srcNode.heading.IsOpposite(heading) {
				continue
			}
			if neighbor.TileType() == NodeType {
				tgtNode := neighbor.(*NodeTile)
				if tgtNode.heading != Undefined {
					// create new node if getting to a node from a different heading.
					// Nodes have different cost depending whether they were approached straight ahead
					// or from the side
					// case: >o> -  o cost is 1
					// case: o>
					//       ^   -  o costs 1000
					if tgtNode.heading != heading {
						m.tiles[NewPosition(tgtNode.pos.row, tgtNode.pos.col, heading)] = NewNodeTile(tgtNode.pos.row, tgtNode.pos.col)
						tgtNode = m.tiles[NewPosition(tgtNode.pos.row, tgtNode.pos.col, heading)].(*NodeTile)
					}
				}
				cost := srcNode.cost + srcNode.heading.Cost(heading) + 1
				tgtCost := tgtNode.cost
				if cost < tgtCost {
					tgtNode.cost = cost
					tgtNode.preTile = []*NodeTile{srcNode}
					tgtNode.heading = heading
					heap.Push(pq, tgtNode)
				}
				if cost == tgtCost {
					tgtNode.preTile = append(tgtNode.preTile, srcNode)
					tgtNode.heading = heading
					heap.Push(pq, tgtNode)	// ????		
				}
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

func (m Maze) NeighborTile(tile NodeTile, h Heading) Tile {
	switch h {
	case North:
		return m.Tile(tile.pos.row-1, tile.pos.col, h)
	case East:
		return m.Tile(tile.pos.row, tile.pos.col+1, h)
	case South:
		return m.Tile(tile.pos.row+1, tile.pos.col, h)
	case West:
		return m.Tile(tile.pos.row, tile.pos.col-1, h)
	default:
		panic("Switch exhausted")
	}
}

func (m Maze) Tile(row int, col int, h Heading) Tile {
	pos := NewPosition(row, col, h)
	tile, exists := m.tiles[pos]
	if exists {
		return tile
	} else {
		pos = NewPosition(row, col, Undefined)
		tile, exists = m.tiles[pos]
		if exists {
			return tile
		}
	}
	panic(fmt.Sprintf("Tile at %s does not exist", pos))
}

func (m Maze) AnyTile(row int, col int) Tile {
	var tile Tile
	count := 0
	for _, h := range Headings() {
		if t, exists := m.tiles[NewPosition(row, col, h)]; exists {
			tile = t
			count++
		}
	}
	switch count {
	case 0:
		return m.Tile(row, col, Undefined)
	case 1:
		return tile
	default:
		nt := tile.(*NodeTile)
		nt.heading = AnyHeading
		return nt
	}
}

func (m Maze) String() string {
	var s strings.Builder
	for row := 0; row < m.dimensions.rows; row++ {
		s.WriteString(fmt.Sprintf("%4d. ", row))
		for col := 0; col < m.dimensions.cols; col++ {
			tile := m.AnyTile(row, col)
			s.WriteString(fmt.Sprint(tile))
		}
		s.WriteString(fmt.Sprintln())
	}
    s.WriteString(fmt.Sprintf("Number of Tiles: %d\n", len(m.tiles)))
	return s.String()
}

func (m Maze) ShortestPaths() [][]*NodeTile {
	return m.shortestPaths
}

func (m Maze) PrintPath(path []*NodeTile) string {
	var s strings.Builder
	sp := NewShortestPath(path)
	for row := 0; row < m.dimensions.rows; row++ {
		s.WriteString(fmt.Sprintf("%4d. ", row))
		for col := 0; col < m.dimensions.cols; col++ {
			tile := m.AnyTile(row, col)
			switch ntile := tile.(type) {
			case WallTile:
				s.WriteString(fmt.Sprint(tile))
			case *NodeTile:
				if _, exists := sp.pathByPos[ntile.pos]; exists || ntile.TileType() == FinishType {
					s.WriteString(fmt.Sprint(tile))
				} else {
					s.WriteString(fmt.Sprint(NodeType))
				}
			}
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
				newNode := NewNodeTile(rowno, colno)
				newNode.nodeType = Start
				newNode.heading = East
				newNode.cost = 0
				tiles[NewPosition(rowno, colno, Undefined)] = newNode
				newMaze.startTile = newNode
			case FinishType:
				newNode := NewNodeTile(rowno, colno)
				newNode.nodeType = Finish
				tiles[NewPosition(rowno, colno, Undefined)] = newNode
				newMaze.finishTile = newNode
			case NodeType:
				tiles[NewPosition(rowno, colno, Undefined)] = NewNodeTile(rowno, colno)
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
