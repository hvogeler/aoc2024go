package main

import (
	// "fmt"
	"reflect"
	"testing"
)

func Test_HikingMap(t *testing.T) {

	testData1 := `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`

	t.Run("FromStr", func(t *testing.T) {
		hikingMap := FromStr(&testData1)
		if hikingMap.dimensions != (Dimensions { 8, 8 }) {
			t.Errorf("Expected 7 rows and 8 cols, got %v", hikingMap.dimensions)
		}

		if !reflect.DeepEqual(hikingMap.grid[1], []int{7,8,1,2,1,8,7,4}) {
			t.Errorf("second row incorrect")
		}

		if hikingMap.grid[2][3] != 3 {
			t.Errorf("Expected 3, got %d", hikingMap.grid[2][3])
		}

		if len(hikingMap.TrailHeads()) != 9 {
			t.Errorf("Expected 9 trail heads, got %d", hikingMap.TrailHeads())
		}
	})



}
