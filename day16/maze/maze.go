package maze

import (
	"bufio"
	"fmt"
	"math"
	"strings"
)

type Maze struct {
	tiles          [][]Tile
	nextReindeerId int
	lowScore       int
	reindeers      map[int]Reindeer
	dimensions     Dimensions
}

func (m *Maze) NextReindeerId() int {
	id := m.nextReindeerId
	m.nextReindeerId++
	return id
}

func (m Maze) NeighborTile(pos Position, direction DirectionType) *Tile {
	neighborPos := pos
	switch direction {
	case Up:
		neighborPos.row--
	case Down:
		neighborPos.row++
	case Left:
		neighborPos.col--
	case Right:
		neighborPos.col++
	default:
		panic("Switch exhausted")
	}

	return m.Tile(neighborPos.Coords())
}

func (m Maze) Tile(row, col int) *Tile {
	if !NewPosition(row, col).IsInside(m.dimensions) {
		return nil
	}
	return &m.tiles[row][col]
}

func (m Maze) String() string {
	var s strings.Builder
	var colno int
	var tile Tile
	for rowno, row := range m.tiles {
		s.WriteString(fmt.Sprintf("%4d. ", rowno))
		for colno, tile = range row {
			s.WriteString(fmt.Sprint(tile))
		}
		s.WriteString(fmt.Sprintln())
	}
	s.WriteString(fmt.Sprintf("      0 - %d columns\n", colno))
	for _, r := range m.reindeers {
		s.WriteString(r.String())
	}
	s.WriteString(fmt.Sprintf("Low Score is %d\n", m.lowScore))
	return s.String()
}

func (m *Maze) CloneReindeersAt(pos Position, currentHeading DirectionType) []Reindeer {
	newReindeers := []Reindeer{}
    thisTile := m.Tile(pos.Coords())
	for _, direction := range DirectionTypes() {
		if direction != currentHeading {
			if (*m.NeighborTile(pos, direction)).TileType() != WallType {
				r := Reindeer{
					id:       m.NextReindeerId(),
					heading:  direction,
					position: pos,
				}
                newReindeers = append(newReindeers, r)
                switch (*thisTile).TileType() {
                case UnusedType:
                    *thisTile = TrackMark{
                        reindeerId: r.id,
                        direction: direction,
                        score: 1000,
                    }
                case TrackMarkType:
                    tm, _ := (*thisTile).(TrackMark)
                    tm.reindeerId = r.id
                    tm.direction = direction
                    tm.score += 1000
                } 
			}
		}
	}
    return newReindeers
}

func MazeFromStr(s string) Maze {
	tiles := [][]Tile{}
	scanner := bufio.NewScanner(strings.NewReader(s))
	var firstReeindeerPosition Position
	var cols int
	var rowno int
	for rowno = 0; scanner.Scan(); rowno++ {
		line := scanner.Text()
		if cols == 0 {
			cols = len(line)
		}
		row := make([]Tile, len(line))
		for colno, tileChar := range strings.Split(line, "") {
			switch TileType(tileChar) {
			case WallType:
				row[colno] = Wall{}
			case StartType:
				row[colno] = Start{}
				firstReeindeerPosition = NewPosition(rowno, colno)
			case FinishType:
				row[colno] = Finish{}
			case UnusedType:
				row[colno] = Unused{}
			}
		}
		tiles = append(tiles, row)
	}
	newMaze := Maze{
		tiles:          tiles,
		nextReindeerId: 0,
		lowScore:       math.MaxInt,
		reindeers:      make(map[int]Reindeer),
		dimensions: Dimensions{
			rows: rowno + 1,
			cols: cols}}
	firstReindeer := Reindeer{
		id:       newMaze.NextReindeerId(),
		position: firstReeindeerPosition,
		heading:  Up}
	newMaze.reindeers[firstReindeer.id] = firstReindeer
	return newMaze
}
