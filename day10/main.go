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



func (hMap HikingMap) TrailHeads() []Location {
	trailHeads := []Location{}
	for row := 0; row < hMap.dimensions.rows; row++ {
		for col := 0; col < hMap.dimensions.cols; col++ {
			if hMap.At(row, col) == 0 {
				trailHeads = append(trailHeads, Location{row, col})
			}
		}
	}
	return trailHeads
}

func (hMap HikingMap) At(row int, col int) int {
	return hMap.grid[row][col]
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
