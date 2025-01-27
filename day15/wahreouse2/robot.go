package warehouse2

import (
	wh1 "day15/warehouse"
)

type Robot struct {
	position wh1.Location
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

func (robot Robot) PositionLeft() wh1.Location {
	return robot.position
}

func (robot Robot) PositionRight() wh1.Location {
	return robot.position
}

func (robot Robot) ItemAt(pos wh1.Location) ItemPart {
	if robot.position == pos {
		return Left
	}
	return None
}

func (robot *Robot) SetPosition(loc wh1.Location) {
    robot.position = loc
}

func NewRobot(x, y int) Robot {
	return Robot{position: wh1.NewLocation(x, y)}
}
