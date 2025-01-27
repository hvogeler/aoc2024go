package warehouse2

import (
	wh "day15/warehouse"
)

type Robot struct {
	position wh.Location
}

func (robot Robot) Item() ItemType {
	return RobotItem
}

func (robot Robot) String() string {
	return string(RobotItem)
}

func (robot Robot) Length() int {
	return 1
}

func (robot Robot) PositionLeft() wh.Location {
	return robot.position
}

func (robot Robot) PositionRight() wh.Location {
	return robot.position
}

func (robot Robot) ItemAt(pos wh.Location) ItemPart {
	if robot.position == pos {
		return Left
	}
	return None
}

func NewRobot(x, y int) Robot {
	return Robot{position: wh.NewLocation(x, y)}
}
