package main

type Antenna struct {
    frequency rune
    location Location
}

type GeoLine struct {
    a Antenna
    b Antenna
}

func (line GeoLine) Distance () Distance{
	return Distance {
		line.a.location.row - line.b.location.row,
		line.a.location.col - line.b.location.col,
	}
}

type Distance struct {
    rows int
    cols int
}

type Location struct {
    row int
    col int
}

type MapDimensions struct {
    rows int
    cols int
}