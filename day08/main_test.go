package main

import (
	"fmt"
	"testing"
)

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
	testdata2 := `T....#....
...T......
.T....#...
.........#
..#.......
..........
...#......
..........
....#.....
..........
`

	t.Run("createAntinodesPart2", func(t *testing.T) {
		cityMap := CityMapFromStr(testdata2)
		cityMap.CreateAntinodes2()

		antinodes := cityMap.antinodes()
		if len(antinodes) != 9 {
			t.Errorf("Expected 9 Antinodes, got %d: %v", len(antinodes), antinodes)
		}

		expectedAntinodes := []Location{{0, 5}, {3, 9}, {4, 2}, {8, 4}}
		for _, expectedAntinode := range expectedAntinodes {
			if !cityMap.isAntinode(&expectedAntinode) {
				t.Errorf("Expected %v to be an Antinode, but it is not", expectedAntinode)
			}
		}

		cityMap = CityMapFromStr(testdata)
		cityMap.CreateAntinodes2()

		antinodes = cityMap.antinodes()
		if len(antinodes) != 34 {
			t.Errorf("Expected 34 Antinodes, got %d: %v", len(antinodes), antinodes)
		}

	})

	t.Run("createCityMap", func(t *testing.T) {
		cityMap := CityMapFromStr(testdata)
		if cityMap.dimensions.rows != 12 {
			t.Errorf("Rows wrong. Expected %d, got %d", 12, cityMap.dimensions.rows)
		}
		if cityMap.dimensions.cols != 12 {
			t.Errorf("Cols wrong. Expected %d, got %d", 12, cityMap.dimensions.cols)
		}

		if len(cityMap.antennasByFrequency) != 2 {
			t.Errorf("length of antennasByFrequency should be %d, but is %d", 2, len(cityMap.antennasByFrequency))
		}

		if len(cityMap.antennasByFrequency['A']) != 3 {
			t.Errorf("Expected 3 antennas A, but got %d", len(cityMap.antennasByFrequency['A']))
		}
	})

	t.Run("createLines", func(t *testing.T) {
		cityMap := CityMapFromStr(testdata)
		cityMap.createLines()
		if len(cityMap.linesByAntenna) != 5 {
			t.Errorf("length of linesByAntenna should be 5, but is %d", len(cityMap.linesByAntenna))
		}
		cityMap.linesByAntenna = nil
		cityMap.CreateAntinodes()
		fmt.Printf("Number of Antinodes: %d\n", len(cityMap.antinodes()))
		fmt.Println(cityMap.antinodes())
		if len(cityMap.antinodes()) != 14 {
			t.Errorf("Expected 14 Antinodes, but got %d", len(cityMap.antinodes()))
		}
	})
}

