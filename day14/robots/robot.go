package robots

type Robot struct {
	position Location
	velocity Velocity
}

func NewRobot(loc Location, v Velocity) Robot {
	return Robot{loc, v}
}


