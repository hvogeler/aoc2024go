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
		if len(space.robots) != 12 {
			t.Errorf("Expected 11 robots. got %d", len(space.robots))
		}
		robot2 := space.robots[1]
		if robot2.position != (Location{6, 3}) {
			t.Errorf("Expected location (6, 3), got %s", robot2.position)
		}
		if robot2.velocity != (Velocity{-1, -3}) {
			t.Errorf("Expected location (-1, -3), got %s", robot2.velocity)
		}
	})
}
