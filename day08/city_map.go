package main

import (
	"bufio"
	"fmt"
	"strings"
)

const DOT = '.'
const ANTINODE = '#'

type CityMap struct {
	antennas            []Antenna
	dimensions          MapDimensions
	antennasByFrequency map[rune][]Antenna
	linesByAntenna      map[Antenna][]GeoLine
	antinodesByLine     map[GeoLine][]Location
}

func (cityMap CityMap) isLocationWithinCity(location Location) bool {
	if location.row < 0 || location.row >= cityMap.dimensions.rows {
		return false
	}
	if location.col < 0 || location.col >= cityMap.dimensions.cols {
		return false
	}
	return true
}

func (cityMap CityMap) antinodes() []Location {
	antinodeSet := make(map[Location]bool)
	for _, location := range flattenArray(getMapValues(cityMap.antinodesByLine)) {
		antinodeSet[location] = true
	}
	return getMapKeys(antinodeSet)
}

func (cityMap CityMap) lines() []GeoLine {
	return flattenArray(getMapValues(cityMap.linesByAntenna))
}

func (cityMap *CityMap) createLines() {
	cityMap.linesByAntenna = make(map[Antenna][]GeoLine)
	for freq := range cityMap.antennasByFrequency {
		antennas := cityMap.antennasByFrequency[freq]
		for i := 0; i < len(antennas)-1; i++ {
			lines := make([]GeoLine, 0)
			for j := i + 1; j < len(antennas); j++ {
				lines = append(lines, GeoLine{antennas[i], antennas[j]})
			}
			cityMap.linesByAntenna[antennas[i]] = lines
		}
	}
}

func (cityMap *CityMap) CreateAntinodes() {
	if cityMap.linesByAntenna == nil {
		cityMap.createLines()
	}
	cityMap.antinodesByLine = make(map[GeoLine][]Location)
	for _, line := range cityMap.lines() {
		d := line.Distance()
		antinodeSet := make(map[Location]bool, 0)
		cityMap.validateAddAntinode(antinodeSet, line, Location{line.a.location.row - d.rows, line.a.location.col - d.cols})
		cityMap.validateAddAntinode(antinodeSet, line, Location{line.a.location.row + d.rows, line.a.location.col + d.cols})
		cityMap.validateAddAntinode(antinodeSet, line, Location{line.b.location.row - d.rows, line.b.location.col - d.cols})
		cityMap.validateAddAntinode(antinodeSet, line, Location{line.b.location.row + d.rows, line.b.location.col + d.cols})
		cityMap.antinodesByLine[line] = getMapKeys(antinodeSet)
	}
}

func (cityMap CityMap) validateAddAntinode(antinodeSet map[Location]bool, line GeoLine, antinodeCandidate Location) {
	if cityMap.isLocationWithinCity(antinodeCandidate) && antinodeCandidate != line.a.location && antinodeCandidate != line.b.location {
		antinodeSet[antinodeCandidate] = true
	}
}

func (cityMap CityMap) getAntennaAt(location Location) *Antenna {
	for _, antenna := range cityMap.antennas {
		if antenna.location == location {
			return &antenna
		}
	}
	return nil
}

func (cityMap CityMap) isAntinodeAt(location Location) bool {
	for _, antinode := range cityMap.antinodes() {
		if antinode == location {
			return true
		}
	}
	return false
}

func (cityMap CityMap) String() string {
	result := ""
	for row := 0; row < cityMap.dimensions.rows; row++ {
		result += fmt.Sprintf("Line %3d: ", row+1)
		for col := 0; col < cityMap.dimensions.cols; col++ {
			objectAtLocation := DOT
			if antenna := cityMap.getAntennaAt(Location{row, col}); antenna != nil {
				objectAtLocation = antenna.frequency
			} else if cityMap.isAntinodeAt(Location{row, col}) {
				objectAtLocation = ANTINODE
			}
			result += string(objectAtLocation)
		}
		result += "\n"
	}
	return result
}

func (cityMap *CityMap) addAntenna(antenna Antenna) {
	cityMap.antennas = append(cityMap.antennas, antenna)
	antennas, exists := cityMap.antennasByFrequency[antenna.frequency]
	if !exists {
		antennas = make([]Antenna, 0)
	}

	if cityMap.antennasByFrequency == nil {
		cityMap.antennasByFrequency = make(map[rune][]Antenna)
	}
	cityMap.antennasByFrequency[antenna.frequency] = append(antennas, antenna)
}

func cityMapFromStr(s string) CityMap {
	var cityMap CityMap
	scanner := bufio.NewScanner(strings.NewReader(s))
	row := 0
	var mapDims MapDimensions
	for scanner.Scan() {
		line := scanner.Text()
		if row == 0 {
			mapDims.cols = len(line)
		} else {
			if len(line) != mapDims.cols {
				panic(fmt.Sprintf("Rows %d has a different number of columns", row))
			}
		}
		runes := []rune(line)
		for col := 0; col < len(runes); col++ {
			currentRune := runes[col]
			if currentRune != '.' && currentRune != '#' {
				cityMap.addAntenna(Antenna{currentRune, Location{row, col}})
			}
		}
		row++
	}
	mapDims.rows = row
	cityMap.dimensions = mapDims
	return cityMap
}
