package maze

import (
	"bufio"
	"fmt"
	"strings"
)

type Maze struct {
	tiles [][]Tile
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
	return s.String()
}

func MazeFromStr(s string) Maze {
	tiles := [][]Tile{}
	scanner := bufio.NewScanner(strings.NewReader(s))
	for rowno := 0; scanner.Scan(); rowno++ {
		line := scanner.Text()
		row := make([]Tile, len(line))
		for colno, tileChar := range strings.Split(line, "") {
			switch TileType(tileChar) {
			case WallType:
				row[colno] = Wall{}
			case StartType:
				row[colno] = Start{}
			case FinishType:
				row[colno] = Finish{}
			case UnusedType:
				row[colno] = Unused{}
			}
		}
		tiles = append(tiles, row)
	}
	return Maze{tiles}
}
