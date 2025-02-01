package maze2

import (
	"bufio"
	"fmt"
	"strings"
)

type Maze struct {
	tiles      map[Position]Tile
	dimensions Dimensions
}

func (m Maze) Tile(row int, col int) *Tile {
	pos := NewPosition(row, col)
	tile, exists := m.tiles[pos]
	if exists {
		return &tile
	}
	panic(fmt.Sprintf("Tile at %s does not exist", pos))
}

func (m Maze) String() string {
	var s strings.Builder
	for row := 0; row < m.dimensions.rows; row++ {
		s.WriteString(fmt.Sprintf("%4d. ", row))
		for col := 0; col < m.dimensions.cols; col++ {
			tile := m.Tile(row, col)
			s.WriteString(fmt.Sprint(*tile))
		}
		s.WriteString(fmt.Sprintln())
	}
	return s.String()
}

func MazeFromStr(s string) Maze {
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
			case Wall:
				tiles[NewPosition(rowno, colno)] = WallTile{}
			case Start:
				tiles[NewPosition(rowno, colno)] = StartTile{}
			case Finish:
				tiles[NewPosition(rowno, colno)] = FinishTile{}
			case Node:
				tiles[NewPosition(rowno, colno)] = NewNodeTile(rowno, colno)
			}
		}
	}
	newMaze := Maze{
		tiles: tiles,
		dimensions: Dimensions{
			rows: rowno,
			cols: cols}}
	return newMaze
}

type Dimensions struct {
	rows int
	cols int
}
