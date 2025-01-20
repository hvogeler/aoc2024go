package robots

import (
	"fmt"
	"testing"
)


func Test_space(t *testing.T) {
	t.Run("Qaudrants", func(t *testing.T) {
		data := ReadData("../example.dat")
		fmt.Println(data)
		space := SpaceFromString(data, Dimensions{11, 7})
		fmt.Println(space)
		if len(space.robots) != 12 {
			t.Errorf("Expected 11 robots. got %d", len(space.robots))
		}
		loc, dim := space.QuadrantCoords(topLeft)
		if loc != (Location{0, 0} ) || dim != (Dimensions{5, 3}){
			t.Errorf("Quadrant 1 wrong")
		}
		loc, dim = space.QuadrantCoords(topRight)
		if loc != (Location{6, 0} ) || dim != (Dimensions{5, 3}){
			t.Errorf("Quadrant 2 wrong")
		}
		loc, dim = space.QuadrantCoords(bottomLeft)
		if loc != (Location{0, 4} ) || dim != (Dimensions{5, 3}){
			t.Errorf("Quadrant 4 wrong")
		}
		loc, dim = space.QuadrantCoords(bottomRight)
		if loc != (Location{6, 4} ) || dim != (Dimensions{5, 3}){
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

		if space.CountAllQuadrants() != 8 {
			t.Errorf("Expected 8 robots in all quadrants, got %d", space.CountAllQuadrants())
		}
	})

	
}
