package robots

import (
	"fmt"
	"testing"
)


func Test_robot(t *testing.T) {
	t.Run("good1", func(t *testing.T) {
		data := ReadData("../example.dat")
		fmt.Println(data)
		space := SpaceFromString(data, Dimensions{11, 7})
		fmt.Println(space)
		robot2 := space.robots[1]
		if robot2.tile.location != (Location{6, 3}) {
			t.Errorf("Expected location (6, 3), got %s", robot2.tile.location)
		}
		if robot2.velocity != (Velocity{-1, -3}) {
			t.Errorf("Expected location (-1, -3), got %s", robot2.velocity)
		}
		
		tile := space.Tile(6, 3)
		robotptr := tile.robots[0]
		fmt.Printf("%p, %p\n", robotptr, robot2)
		if robotptr != robot2 {
			fmt.Printf("%p, %p",robotptr, robot2)
			t.Errorf("robot pointer in tile is wrong")
		}

		if space.Tile(3, 0).countRobots() != 2 {
			t.Errorf("Expected 2 robots on tile (3,0), got %d", space.Tile(3, 0).countRobots())
		}
	})


}
