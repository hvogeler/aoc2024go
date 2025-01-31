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
	reindeers      []Reindeer
	dimensions     Dimensions
}

func (m *Maze) MoveReindeer() {
	for i, reindeer := range m.reindeers {
		if !reindeer.IsAlive() {
			continue
		}
		r := &m.reindeers[i]
		if reindeer.score >= m.lowScore {
			r.Kill("Current low score exceeded")
			continue
		}
		nextTile, nextTilePos := m.NeighborTile(reindeer.Position(), reindeer.heading)
		switch (*nextTile).TileType() {
		case FinishType:
			r.score++
			r.SetPosition(nextTilePos)
			if r.score < m.lowScore {
				m.lowScore = r.score
			}
			r.Kill("End of maze reached")
		case UnusedType:
			r.score++
			r.SetPosition(nextTilePos)
			*nextTile = TrackMark{
				reindeerId: r.id,
				heading:    r.heading,
				score:      r.score,
			}
		case TrackMarkType:
			tm := (*nextTile).(TrackMark)
			if tm.heading.TurnRate(reindeer.heading) == 0 {
                if r.AlreadyVisited(nextTilePos) {
                    r.Kill("Running in a loop")
                }
				r.score++
				if r.score < tm.score {
					r.SetPosition(nextTilePos)
					tm.reindeerId = r.id
					tm.score = r.score
					tm.heading = r.heading
					m.tiles[nextTilePos.row][nextTilePos.col] = tm
				} else {
					r.Kill("Current Reindeer score exceeds latest score on that tile")
				}
			}
			if tm.heading.TurnRate(reindeer.heading) == 180 {
				r.Kill("Reversing Heading")
			}
			if tm.heading.TurnRate(reindeer.heading) == 90 {
				r.score++
				r.SetPosition(nextTilePos)
			}
			// case WallType:
			// 	m.CloneReindeer(&m.reindeers[i])
		}
		if m.reindeers[i].IsAlive() {
			m.CloneReindeer(&m.reindeers[i])
		}

	}
}

// Return number of cloned reindeers. 0 means dead end. 1 means just a turn
func (m *Maze) CloneReindeer(r1 *Reindeer) int {
	if m.IsDeadEnd(r1.Position(), r1.heading) {
		r1.Kill("Dead end: TODO check if this test can/must be removed")
		return 0
	}
	pos := r1.Position()
	currentHeading := r1.heading
	thisTile := m.Tile(pos.Coords())
	neighborTileCurrentHeading, _ := (*m).NeighborTile(pos, currentHeading)
	cloneCount := 0
	for _, newHeading := range HeadingTypes() {
		if newHeading.TurnRate(currentHeading) != 0 {
			neighborTile, _ := (*m).NeighborTile(pos, newHeading)
			if (*neighborTile).TileType() != WallType {
				if (*neighborTileCurrentHeading).TileType() == WallType {
					r1.Kill("Wall straight ahead. Kill this Reindeer and send a clone in the new heading(s)")
					// continue
				}
				if newHeading.TurnRate(currentHeading) == 180 {
					continue
				}
				score := currentHeading.Score(newHeading)
				rClone := r1.Clone(m.NextReindeerId(), newHeading, score)
				if _, exists := m.ReindeerById(rClone.id); exists {
					panic(fmt.Sprintf("Duplicate Reindeer ID: %d", rClone.id))
				}
				m.reindeers = append(m.reindeers, rClone)
				cloneCount++
				switch (*thisTile).TileType() {
				case UnusedType:
					*thisTile = TrackMark{
						reindeerId: rClone.id,
						heading:    newHeading,
						score:      score,
					}
				case TrackMarkType:
					tm, _ := (*thisTile).(TrackMark)
					*thisTile = TrackMark{
						reindeerId: rClone.id,
						heading:    newHeading,
						score:      tm.score + score,
					}
				}
			}
		}
	}
	return cloneCount
}

func (m Maze) CountAlive() int {
	sum := 0
	for _, r := range m.reindeers {
		if r.IsAlive() {
			sum++
		}
	}
	return sum
}

func (m *Maze) NextReindeerId() int {
	id := m.nextReindeerId
	m.nextReindeerId++
	return id
}

func (m Maze) NeighborTile(pos Position, heading HeadingType) (*Tile, Position) {
	neighborPos := pos
	switch heading {
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

	return m.Tile(neighborPos.Coords()), neighborPos
}

func (m Maze) IsDeadEnd(pos Position, currentHeading HeadingType) bool {
	wallCount := 0
	nTileCurrentHeadingIsWall := false
	nTileOppositeHeadingIsWall := false
	oppTile, _ := m.NeighborTile(pos, currentHeading.OppositeHeading())
	if (*oppTile).TileType() == WallType {
		nTileOppositeHeadingIsWall = true
	}
	for _, heading := range HeadingTypes() {
		nTile, _ := m.NeighborTile(pos, heading)
		if (*nTile).TileType() == WallType {
			wallCount++
			if heading == currentHeading {
				nTileCurrentHeadingIsWall = true
			}
		}
	}

	return wallCount == 3 && nTileCurrentHeadingIsWall && !nTileOppositeHeadingIsWall
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
		if r.IsAlive() {
			s.WriteString(r.String())
		}
	}
	s.WriteString(fmt.Sprintf("Low Score is %d\n", m.lowScore))
	return s.String()
}

func (m Maze) ReindeerById(id int) (*Reindeer, bool) {
	for _, reindeer := range m.reindeers {
		if reindeer.id == id {
			return &reindeer, true
		}
	}
	return nil, false
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
		reindeers:      []Reindeer{},
		dimensions: Dimensions{
			rows: rowno + 1,
			cols: cols}}
	firstReindeer := NewReindeer(newMaze.NextReindeerId(), firstReeindeerPosition)
	newMaze.reindeers = append(newMaze.reindeers, firstReindeer)
	newMaze.CloneReindeer(&newMaze.reindeers[0])
	return newMaze
}
