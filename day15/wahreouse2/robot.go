package warehouse2

import (
	wh "day15/warehouse"
)

type Robot struct {
	position wh.Location
}

func NewRobot(x, y int) Robot {
	return Robot{position: wh.NewLocation(x, y)}
}
