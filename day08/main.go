package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	bytes, err := os.ReadFile("example.dat")
    rawMap := string(bytes)
    if err != nil {
        panic(err)
    }

    cityMap := cityMapFromStr(rawMap)
    fmt.Println(cityMap)
}

const DOT = '.'
const ANTINODE = '#'

type CityMap struct {
    antennas []Antenna
    dimensions MapDimensions
    antennasByFrequency map[rune][]Antenna
    linesByAntenna map[Antenna][]GeoLine
}

func (cityMap *CityMap) createLines() {
    cityMap.linesByAntenna = make(map[Antenna][]GeoLine)
    for freq := range cityMap.antennasByFrequency {
        antennas := cityMap.antennasByFrequency[freq]
        for i := 0; i < len(antennas) - 1; i++ {
            lines := make([]GeoLine, 0)
            for j := i + 1; j < len(antennas); j++ {
                lines = append(lines, GeoLine { antennas[i], antennas[j]})
            }
            cityMap.linesByAntenna[antennas[i]] = lines
        } 
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

func (cityMap CityMap) String() string {
    result := ""
    for row := 0; row < cityMap.dimensions.rows; row++ {
        result += fmt.Sprintf("Line %3d: ", row + 1)
        for col := 0; col < cityMap.dimensions.cols; col++ {
            if antenna := cityMap.getAntennaAt(Location { row, col }); antenna != nil {
                result += string(antenna.frequency)
            } else {
                result += string(DOT)
            }
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
                cityMap.addAntenna(Antenna{ currentRune, Location{ row, col }})
            } 
        }
        row++
    }
    mapDims.rows = row
    cityMap.dimensions = mapDims
    return cityMap
}

type Antenna struct {
    frequency rune
    location Location
}

type GeoLine struct {
    a Antenna
    b Antenna
}

type Distance struct {
    rows int64
    cols int64
}

type Location struct {
    row int
    col int
}

type MapDimensions struct {
    rows int
    cols int
}