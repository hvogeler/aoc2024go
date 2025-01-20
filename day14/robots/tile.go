package robots

type Tile struct {
	location Location
	robots []*Robot
}

func (tile *Tile) AddRobot(robot *Robot) {
	tile.robots = append(tile.robots, robot)
}

func (tile Tile) countRobots() int {
	return len(tile.robots)
}