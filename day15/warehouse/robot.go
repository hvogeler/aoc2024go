package warehouse

const RobotRune rune = '@'

type Robot struct {
	position Location
}

func NewRobot(x, y int) Robot {
	return Robot{position: NewLocation(x, y)}
}
