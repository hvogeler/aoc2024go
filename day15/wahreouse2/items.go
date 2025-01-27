package warehouse2

import (
	wh1 "day15/warehouse"
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

type Item interface {
	Item() ItemType
	PositionRight() wh1.Location
	PositionLeft() wh1.Location
	ItemAt(pos wh1.Location) ItemPart
	Length() int
	String() string
    SetPosition(loc wh1.Location)
}

type ItemPart int

const (
	None ItemPart = iota
	Left
	Right
)

type Wall struct {
	position wh1.Location
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

func (wall Wall) PositionLeft() wh1.Location {
	return wall.position
}

func (wall Wall) PositionRight() wh1.Location {
	return wh1.NewLocation(wall.position.X()+1, wall.position.Y())
}

func (wall Wall) ItemAt(pos wh1.Location) ItemPart {
	switch pos {
	case wall.PositionLeft():
		return Left
	case wall.PositionRight():
		return Right
	default:
		return None
	}
}

func (wall *Wall) SetPosition(loc wh1.Location) {
    wall.position = loc
}

func NewWall(leftx int, lefty int) Wall {
	wall := new(Wall)
	wall.position = wh1.NewLocation(leftx, lefty)
	return *wall
}

type Box struct {
	position wh1.Location
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

func (box Box) PositionLeft() wh1.Location {
	return box.position
}

func (box Box) PositionRight() wh1.Location {
	return wh1.NewLocation(box.position.X()+1, box.position.Y())
}

func (box Box) ItemAt(pos wh1.Location) ItemPart {
	switch pos {
	case box.PositionLeft():
		return Left
	case box.PositionRight():
		return Right
	default:
		return None
	}
}

func (box *Box) SetPosition(loc wh1.Location) {
    box.position = loc
}

func NewBox(leftx int, lefty int) Box {
	box := new(Box)
	box.position = wh1.NewLocation(leftx, lefty)
	return *box
}
