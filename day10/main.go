package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	bytes, err := os.ReadFile("testdata.dat")
	if err != nil {
		panic(err)
	}
	data := string(bytes)
	hikingMap := FromStr(&data)

	sumAll := 0
	for _, startLocation := range hikingMap.TrailHeads() {
		sum9 := hikingMap.Walk(startLocation)
		fmt.Printf("Start: %s  -  Sum = %d\n", startLocation, sum9)
		sumAll += sum9
	}
	fmt.Println("Sum All = ", sumAll)
	if sumAll != 459 {
		panic("Wrong result for Part 1")
	}

	// Part2
	sumAll = 0
	for _, startLocation := range hikingMap.TrailHeads() {
		sum9 := hikingMap.Walk2(startLocation)
		fmt.Printf("Start: %s  -  Sum = %d\n", startLocation, sum9)
		sumAll += sum9
	}
	fmt.Println("Sum All = ", sumAll)
	if sumAll != 1034 {
		panic("Wrong result for Part 2")
	}

}

type HikingMap struct {
	raw        string
	grid       [][]int
	dimensions Dimensions
}

// Walk the trails and sum the number of peaks the reach
func (hMap HikingMap) Walk2(start Location) int {

	return hMap.stepFrom2(start, 0)
}

func (hMap HikingMap) stepFrom2(loc Location, level int) int {
	if hMap.At(loc) == 9 {
		return 1
	}

	// Check for 1 increase in every direction
	sum_9_heights := 0
	if hMap.At(loc) == hMap.up(loc) - 1 {
		sum_9_heights += hMap.stepFrom2(loc.Up(), level + 1)
	}
	if hMap.At(loc) == hMap.down(loc) - 1 {
		sum_9_heights += hMap.stepFrom2(loc.Down(), level + 1)
	}
	if hMap.At(loc) == hMap.left(loc) - 1 {
		sum_9_heights += hMap.stepFrom2(loc.Left(), level + 1)
	}
	if hMap.At(loc) == hMap.right(loc) - 1 {
		sum_9_heights += hMap.stepFrom2(loc.Right(), level + 1)
	}
	fmt.Printf("Unwind Recursion Level: %d\n", level)
	return sum_9_heights
}

func (hMap HikingMap) Walk(start Location) int {
	hMap.stepFrom(start, 0)
	return len(hMap.stepFrom(start, 0))
}
func (hMap HikingMap) stepFrom(loc Location, level int) []Location {
	if hMap.At(loc) == 9 {
		return []Location{loc}
	}

	// Check for 1 increase in every direction
	allTargets := make(map[Location]bool)
	if hMap.At(loc) == hMap.up(loc)-1 {
		targets := hMap.stepFrom(loc.Up(), level+1)
		for _, target := range targets {
			allTargets[target] = true
		}
	}
	if hMap.At(loc) == hMap.down(loc)-1 {
		targets := hMap.stepFrom(loc.Down(), level+1)
		for _, target := range targets {
			allTargets[target] = true
		}
	}
	if hMap.At(loc) == hMap.left(loc)-1 {
		targets := hMap.stepFrom(loc.Left(), level+1)
		for _, target := range targets {
			allTargets[target] = true
		}
	}
	if hMap.At(loc) == hMap.right(loc)-1 {
		targets := hMap.stepFrom(loc.Right(), level+1)
		for _, target := range targets {
			allTargets[target] = true
		}
	}

	uniqueTargets := []Location{}
	for target := range allTargets {
		uniqueTargets = append(uniqueTargets, target)
	}

	fmt.Printf("Unwind Recursion Level: %d - %v\n", level, uniqueTargets)
	return uniqueTargets
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
	return Location{loc.row - 1, loc.col}
}

func (loc Location) Down() Location {
	return Location{loc.row + 1, loc.col}
}

func (loc Location) Left() Location {
	return Location{loc.row, loc.col - 1}
}

func (loc Location) Right() Location {
	return Location{loc.row, loc.col + 1}
}

func (loc Location) String() string {
	return fmt.Sprintf("(%d, %d)", loc.row, loc.col)
}
