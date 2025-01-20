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
	})

	
}
