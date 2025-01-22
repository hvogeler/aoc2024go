package warehouse

const RobotRune string = "@"

type Robot struct {
	position Location
}

func NewRobot(x, y int) Robot {
	return Robot{position: NewLocation(x, y)}
}
