package robots

type Robot struct {
	id int
	velocity Velocity
	tile *Tile
}

func NewRobot(id int, v Velocity) *Robot {
	return &Robot{id, v, nil}
}


