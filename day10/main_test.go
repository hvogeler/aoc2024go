package main

import (
	// "fmt"
	"fmt"
	"reflect"
	"testing"
)

func Test_HikingMap(t *testing.T) {

	// Part 2 tests
	t.Run("stepFrom2a", func(t *testing.T) {
		testData := `.....0.
..4321.
..5..2.
..6543.
..7..4.
..8765.
..9....`
		hikingMap := FromStr(&testData)
		if hikingMap.dimensions != (Dimensions{7, 7}) {
			t.Errorf("Expected 7 rows and 7 cols, got %v", hikingMap.dimensions)
		}

		startLocation := hikingMap.TrailHeads()[0]
		if startLocation != (Location{0, 5}) {
			t.Errorf("Expected Trailhead at (0, 5), got %s", startLocation)
		}

		sum9 := hikingMap.Walk2(startLocation)
		fmt.Println("Sum = ", sum9)
		if sum9 != 3 {
			t.Errorf("Expected sum9 to be 3, got %d", sum9)
		}
	})

	t.Run("stepFrom2b", func(t *testing.T) {
		testData := `..90..9
...1.98
...2..7
6543456
765.987
876....
987....`
		hikingMap := FromStr(&testData)
		if hikingMap.dimensions != (Dimensions{7, 7}) {
			t.Errorf("Expected 7 rows and 7 cols, got %v", hikingMap.dimensions)
		}

		sumAll := 0
		for _, startLocation := range hikingMap.TrailHeads() {
			sum9 := hikingMap.Walk2(startLocation)
			fmt.Printf("Start: %s  -  Sum = %d\n", startLocation, sum9)
			sumAll += sum9
		}
		fmt.Println("Sum All = ", sumAll)
		if sumAll != 13 {
			t.Errorf("Expected sum to be 13, got %d", sumAll)
		}

	})

	t.Run("stepFrom2c", func(t *testing.T) {
		testData := `012345
123456
234567
345678
4.6789
56789.`
		hikingMap := FromStr(&testData)
		if hikingMap.dimensions != (Dimensions{6, 6}) {
			t.Errorf("Expected 6 rows and 6 cols, got %v", hikingMap.dimensions)
		}

		sumAll := 0
		for _, startLocation := range hikingMap.TrailHeads() {
			sum9 := hikingMap.Walk2(startLocation)
			fmt.Printf("Start: %s  -  Sum = %d\n", startLocation, sum9)
			sumAll += sum9
		}
		fmt.Println("Sum All = ", sumAll)
		if sumAll != 227 {
			t.Errorf("Expected sum to be 227, got %d", sumAll)
		}

	})

	t.Run("stepFrom2d", func(t *testing.T) {

		testData := `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`

		hikingMap := FromStr(&testData)
		if hikingMap.dimensions != (Dimensions{8, 8}) {
			t.Errorf("Expected 8 rows and 8 cols, got %v", hikingMap.dimensions)
		}

		sumAll := 0
		for _, startLocation := range hikingMap.TrailHeads() {
			sum9 := hikingMap.Walk2(startLocation)
			fmt.Printf("Start: %s  -  Sum = %d\n", startLocation, sum9)
			sumAll += sum9
		}
		fmt.Println("Sum All = ", sumAll)
		if sumAll != 81 {
			t.Errorf("Expected sum to be 81, got %d", sumAll)
		}

	})

	// Part 1 tests
	t.Run("stepFromDupTarget", func(t *testing.T) {
		testData := `...0...
...1...
...2345
6543456
7.....7
8.....8
9.....9`
		hikingMap := FromStr(&testData)
		if hikingMap.dimensions != (Dimensions{7, 7}) {
			t.Errorf("Expected 7 rows and 7 cols, got %v", hikingMap.dimensions)
		}

		startLocation := hikingMap.TrailHeads()[0]
		if startLocation != (Location{0, 3}) {
			t.Errorf("Expected Trailhead at (0, 3), got %s", startLocation)
		}

		sum9 := hikingMap.Walk(startLocation)
		fmt.Println("Sum = ", sum9)
		if sum9 != 2 {
			t.Errorf("Expected sum9 to be 2, got %d", sum9)
		}
	})

	t.Run("stepFrom1", func(t *testing.T) {
		testData := `...0...
...1...
...2...
6543456
7.....7
8.....8
9.....9`
		hikingMap := FromStr(&testData)
		if hikingMap.dimensions != (Dimensions{7, 7}) {
			t.Errorf("Expected 7 rows and 7 cols, got %v", hikingMap.dimensions)
		}

		startLocation := hikingMap.TrailHeads()[0]
		if startLocation != (Location{0, 3}) {
			t.Errorf("Expected Trailhead at (0, 3), got %s", startLocation)
		}

		sum9 := hikingMap.Walk(startLocation)
		fmt.Println("Sum = ", sum9)
		if sum9 != 2 {
			t.Errorf("Expected sum9 to be 2, got %d", sum9)
		}
	})

	t.Run("stepFrom2", func(t *testing.T) {
		testData := `10..9..
2...8..
3...7..
4567654
...8..3
...9..2
.....01`
		hikingMap := FromStr(&testData)
		if hikingMap.dimensions != (Dimensions{7, 7}) {
			t.Errorf("Expected 7 rows and 8 cols, got %v", hikingMap.dimensions)
		}

		startLocation := hikingMap.TrailHeads()[0]
		if startLocation != (Location{0, 1}) {
			t.Errorf("Expected Trailhead at (0, 3), got %s", startLocation)
		}

		sumAll := 0
		for _, startLocation := range hikingMap.TrailHeads() {
			sum9 := hikingMap.Walk(startLocation)
			fmt.Printf("Start: %s  -  Sum = %d\n", startLocation, sum9)
			sumAll += sum9
		}
		fmt.Println("Sum All = ", sumAll)
		if sumAll != 3 {
			t.Errorf("Expected sum to be 3, got %d", sumAll)
		}
	})

	t.Run("FromStr", func(t *testing.T) {

		testData := `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`

		hikingMap := FromStr(&testData)
		if hikingMap.dimensions != (Dimensions{8, 8}) {
			t.Errorf("Expected 8 rows and 8 cols, got %v", hikingMap.dimensions)
		}

		if !reflect.DeepEqual(hikingMap.grid[1], []int{7, 8, 1, 2, 1, 8, 7, 4}) {
			t.Errorf("second row incorrect")
		}

		if hikingMap.grid[2][3] != 3 {
			t.Errorf("Expected 3, got %d", hikingMap.grid[2][3])
		}

		if len(hikingMap.TrailHeads()) != 9 {
			t.Errorf("Expected 9 trail heads, got %d", hikingMap.TrailHeads())
		}

		sumAll := 0
		for _, startLocation := range hikingMap.TrailHeads() {
			sum9 := hikingMap.Walk(startLocation)
			fmt.Printf("Start: %s  -  Sum = %d\n", startLocation, sum9)
			sumAll += sum9
		}
		fmt.Println("Sum All = ", sumAll)
		if sumAll != 36 {
			t.Errorf("Expected sum to be 36, got %d", sumAll)
		}

	})

}
