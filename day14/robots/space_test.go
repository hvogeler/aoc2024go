package robots

import (
	"fmt"
	"testing"
)

func Test_space(t *testing.T) {

	t.Run("Move Robot", func(t *testing.T) {
		data := ReadData("../example.dat")
		space := SpaceFromString(data, Dimensions{11, 7})
		fmt.Println(space)
		space.MoveRobot(0, 1)
		fmt.Println(space)
		if space.Tile(0, 4).CountRobots() != 0 || space.Tile(3, 1).CountRobots() != 1 {
			t.Errorf("Move failed")
		}
	})

	t.Run("Move Robot 10", func(t *testing.T) {
		data := ReadData("../example.dat")
		space := SpaceFromString(data, Dimensions{11, 7})
		fmt.Println(space)
		for i := 0; i < 5; i++ {
			space.MoveRobot(10, 1)
			fmt.Println(space)
		}
		if space.Tile(2, 4).CountRobots() != 0 || space.Tile(1, 3).CountRobots() != 1 {
			t.Errorf("Move failed")
		}
	})

	t.Run("Move all 100", func(t *testing.T) {
		data := ReadData("../example.dat")
		space := SpaceFromString(data, Dimensions{11, 7})
		fmt.Println(space)
		space.MoveRobots(100)
		fmt.Println(space)
		
		if space.SafetyFactor() != 12 {
			t.Errorf("Expected safety factor 12, got %d", space.SafetyFactor())
		}
	})

	t.Run("Move All Robot", func(t *testing.T) {
		data := ReadData("../example.dat")
		space := SpaceFromString(data, Dimensions{11, 7})
		fmt.Println(space)
		space.MoveRobots(1)
		fmt.Println(space)
		if space.Tile(0, 4).CountRobots() != 0 || space.Tile(3, 1).CountRobots() != 1 {
			t.Errorf("Move failed")
		}
	})

	t.Run("Qaudrants", func(t *testing.T) {
		data := ReadData("../example.dat")
		fmt.Println(data)
		space := SpaceFromString(data, Dimensions{11, 7})
		fmt.Println(space)
		if len(space.robots) != 12 {
			t.Errorf("Expected 11 robots. got %d", len(space.robots))
		}
		loc, dim := space.QuadrantCoords(topLeft)
		if loc != (Location{0, 0}) || dim != (Dimensions{5, 3}) {
			t.Errorf("Quadrant 1 wrong")
		}
		loc, dim = space.QuadrantCoords(topRight)
		if loc != (Location{6, 0}) || dim != (Dimensions{5, 3}) {
			t.Errorf("Quadrant 2 wrong")
		}
		loc, dim = space.QuadrantCoords(bottomLeft)
		if loc != (Location{0, 4}) || dim != (Dimensions{5, 3}) {
			t.Errorf("Quadrant 4 wrong")
		}
		loc, dim = space.QuadrantCoords(bottomRight)
		if loc != (Location{6, 4}) || dim != (Dimensions{5, 3}) {
			t.Errorf("Quadrant 4 wrong")
		}

		if space.CountQuadrant(topLeft) != 4 {
			t.Errorf("Expected 4 robot in top left quadrant, got %d", space.CountQuadrant(topLeft))
		}

		if space.CountQuadrant(topRight) != 0 {
			t.Errorf("Expected 0 robot in top right quadrant, got %d", space.CountQuadrant(topRight))
		}

		if space.CountQuadrant(bottomLeft) != 2 {
			t.Errorf("Expected 2 robot in bottom left quadrant, got %d", space.CountQuadrant(bottomLeft))
		}

		if space.CountQuadrant(bottomRight) != 2 {
			t.Errorf("Expected 2 robot in bottom right quadrant, got %d", space.CountQuadrant(bottomRight))
		}

	})

}
