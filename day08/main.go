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

func cityMapFromStr(s string) CityMap {
    var cityMap CityMap
    scanner := bufio.NewScanner(strings.NewReader(s))
    row := 0
    var mapDims MapDimensions
    var antennas []Antenna
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
                antennas = append(antennas, Antenna{ currentRune, Location{ row, col }})
            } 
        }
        row++
    }
    mapDims.rows = row
    cityMap.antennas = antennas
    cityMap.dimensions = mapDims
    return cityMap
}

type Antenna struct {
    frequency rune
    location Location
}

type Location struct {
    row int
    col int
}

type MapDimensions struct {
    rows int
    cols int
}