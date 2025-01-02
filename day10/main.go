package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	bytes, err := os.ReadFile("example1.dat")
	if err != nil {
		panic(err)
	}
	data := string(bytes)
	hMap := FromStr(&data)
	fmt.Println(hMap)

	fmt.Println(hMap.TrailHeads())

}

type HikingMap struct {
	raw        string
	grid       [][]int
	dimensions Dimensions
}

// Walk the trails and sum the number of peaks the reach
func (hMap HikingMap) walk() int {

	return 42
}

func (hMap HikingMap) stepFrom(loc Location) bool {
	if hMap.At(loc) == 9 {
		return true
	}
	if hMap.At(loc) == hMap.up(loc) - 1 {
		hMap.stepFrom(loc.Up())
	}
	if hMap.At(loc) == hMap.down(loc) - 1 {
		hMap.stepFrom(loc.Down())
	}
	if hMap.At(loc) == hMap.left(loc) - 1 {
		hMap.stepFrom(loc.Left())
	}
	if hMap.At(loc) == hMap.right(loc) - 1 {
		hMap.stepFrom(loc.Right())
	}
	return false
}


func (hMap HikingMap) TrailHeads() []Location {
	trailHeads := []Location{}
	for row := 0; row < hMap.dimensions.rows; row++ {
		for col := 0; col < hMap.dimensions.cols; col++ {
			if hMap.At(Location{row, col}) == 0 {
				trailHeads = append(trailHeads, Location{row, col})
			}
		}
	}
	return trailHeads
}

func (hMap HikingMap) At(loc Location) int {
	if loc.row < 0 || loc.row >= hMap.dimensions.rows || loc.col < 0 || loc.col >= hMap.dimensions.cols {
		return -1
	}
	return hMap.grid[loc.row][loc.col]
}

func (hMap HikingMap) up(loc Location) int {
	return hMap.At(loc.Up())
}

func (hMap HikingMap) down(loc Location) int {
	return hMap.At(loc.Down())
}

func (hMap HikingMap) left(loc Location) int {
	return hMap.At(loc.Left())
}

func (hMap HikingMap) right(loc Location) int {
	return hMap.At(loc.Right())
}

func FromStr(s *string) HikingMap {
	hikingMap := new(HikingMap)
	scanner := bufio.NewScanner(strings.NewReader(*s))
	for scanner.Scan() {
		line := scanner.Text()
		row := []int{}
		for _, rune := range line {
			row = append(row, int(rune-0x30))
		}
		hikingMap.grid = append(hikingMap.grid, row)
		if hikingMap.dimensions.cols == 0 {
			hikingMap.dimensions.cols = len(row)
		}
	}
	hikingMap.dimensions.rows = len(hikingMap.grid)
	hikingMap.raw = *s
	return *hikingMap
}

type Dimensions struct {
	rows int
	cols int
}

type Location struct {
	row int
	col int
}

func (loc Location) Up() Location {
	return Location{ loc.row - 1, loc.col }
}

func (loc Location) Down() Location {
	return Location{ loc.row + 1, loc.col }
}

func (loc Location) Left() Location {
	return Location{ loc.row, loc.col -1 }
}

func (loc Location) Right() Location {
	return Location{ loc.row, loc.col + 1 }
}
