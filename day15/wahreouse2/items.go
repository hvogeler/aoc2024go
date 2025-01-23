package warehouse2

import (
	wh "day15/warehouse"
)

const (
	RobotRune  rune = '@'
	WallRune   rune = '#'
	BoxRune    rune = 'O'
	UnusedRune rune = '.'
)

type ItemType string

const (
	RobotItem  ItemType = "@."
	WallItem   ItemType = "##"
	BoxItem    ItemType = "[]"
	UnusedItem ItemType = ".."
)

type Item interface {
	Item() ItemType
	Position() (wh.Location, wh.Location)
}

const (
	Left int = iota
	Right
)

type Wall struct {
	position [2]wh.Location
}

func (wall Wall) Item() ItemType {
	return WallItem
}

func (wall Wall) Position() (wh.Location, wh.Location) {
	return wall.position[0], wall.position[1]
}

func NewWall(leftx int, lefty int) Wall {
	wall := new(Wall)
	wall.position[Left] = wh.NewLocation(leftx, lefty)
	wall.position[Right] = wh.NewLocation(leftx+1, lefty)
	return *wall
}

type Box struct {
	position [2]wh.Location
}

func (box Box) Item() ItemType {
	return BoxItem
}

func (box Box) Position() (wh.Location, wh.Location) {
	return box.position[0], box.position[1]
}

func NewBox(leftx int, lefty int) Box {
	box := new(Box)
	box.position[Left] = wh.NewLocation(leftx, lefty)
	box.position[Right] = wh.NewLocation(leftx+1, lefty)
	return *box
}

