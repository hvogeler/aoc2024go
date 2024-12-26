package main

import "testing"

func Test_CityMap(t *testing.T) {
	testdata := `......#....#
...#....0...
....#0....#.
..#....0....
....0....#..
.#....A.....
...#........
#......#....
........A...
.........A..
..........#.
..........#.
`
	t.Run("createCityMap", func(t *testing.T) {
		cityMap := cityMapFromStr(testdata)
		if cityMap.dimensions.rows != 12 {
			t.Errorf("Rows wrong. Expected %d, got %d", 12, cityMap.dimensions.rows)
		}
		if cityMap.dimensions.cols != 12 {
			t.Errorf("Cols wrong. Expected %d, got %d", 12, cityMap.dimensions.cols)
		}
	})
}