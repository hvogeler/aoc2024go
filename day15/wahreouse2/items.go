package warehouse2

import (
	"strings"
)

const (
	RobotRune  rune = '@'
	WallRune   rune = '#'
	BoxRune    rune = 'O'
	UnusedRune rune = '.'
)

type ItemType string

const (
	RobotItem  ItemType = "@"
	WallItem   ItemType = "##"
	BoxItem    ItemType = "[]"
	UnusedItem ItemType = ".."
)

func (itemt ItemType) String() string {
	var s strings.Builder
	switch itemt {
	case RobotItem:
		s.WriteString("Robot")
	case WallItem:
		s.WriteString("Wall")
	case BoxItem:
		s.WriteString("Box")
	case UnusedItem:
		s.WriteString("Unused")
	}
	return s.String()
}

type Item interface {
	Item() ItemType
	PositionRight() Location
	PositionLeft() Location
	ItemAt(pos Location) ItemPart
	Length() int
	String() string
    SetPosition(loc Location)
}

type ItemPart int

const (
	None ItemPart = iota
	Left
	Right
)

type Wall struct {
	position Location
}

func (wall Wall) Item() ItemType {
	return WallItem
}

func (wall Wall) String() string {
	return string(wall.Item())
}

func (wall Wall) Length() int {
	return 2
}

func (wall Wall) PositionLeft() Location {
	return wall.position
}

func (wall Wall) PositionRight() Location {
	return NewLocation(wall.position.X()+1, wall.position.Y())
}

func (wall Wall) ItemAt(pos Location) ItemPart {
	switch pos {
	case wall.PositionLeft():
		return Left
	case wall.PositionRight():
		return Right
	default:
		return None
	}
}

func (wall *Wall) SetPosition(loc Location) {
    wall.position = loc
}

func NewWall(leftx int, lefty int) Wall {
	wall := new(Wall)
	wall.position = NewLocation(leftx, lefty)
	return *wall
}

type Box struct {
	position Location
}

func (box Box) Item() ItemType {
	return BoxItem
}

func (box Box) String() string {
	return string(box.Item())
}

func (box Box) Length() int {
	return 2
}

func (box Box) PositionLeft() Location {
	return box.position
}

func (box Box) PositionRight() Location {
	return NewLocation(box.position.X()+1, box.position.Y())
}

func (box Box) ItemAt(pos Location) ItemPart {
	switch pos {
	case box.PositionLeft():
		return Left
	case box.PositionRight():
		return Right
	default:
		return None
	}
}

func (box *Box) SetPosition(loc Location) {
    box.position = loc
}

func NewBox(leftx int, lefty int) Box {
	box := new(Box)
	box.position = NewLocation(leftx, lefty)
	return *box
}
