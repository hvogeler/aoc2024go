package warehouse2


type Robot struct {
	position Location
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

func (robot Robot) PositionLeft() Location {
	return robot.position
}

func (robot Robot) PositionRight() Location {
	return robot.position
}

func (robot Robot) ItemAt(pos Location) ItemPart {
	if robot.position == pos {
		return Left
	}
	return None
}

func (robot *Robot) SetPosition(loc Location) {
    robot.position = loc
}

func NewRobot(x, y int) Robot {
	return Robot{position: NewLocation(x, y)}
}
