package robots

type Tile struct {
	location Location
	robots   []*Robot
}

func (tile *Tile) AddRobot(robot *Robot) {
	tile.robots = append(tile.robots, robot)
}

func (tile Tile) CountRobots() int {
	return len(tile.robots)
}

func (tile *Tile) RemoveRobot(robot *Robot) {
	idx, found := tile.FindRobot(robot)
	if found {
		tile.robots = append(tile.robots[:idx], tile.robots[idx+1:]...)
	}
}

func (tile Tile) FindRobot(robotToRemove *Robot) (int, bool) {
	for i, robot := range tile.robots {
		if robot == robotToRemove {
			return i, true
		}
	}
	return -1, false
}
