package warehouse

import (
	"fmt"
	"testing"
)

func Test_xmas_tree(t *testing.T) {

	t.Run("Move Robots", func(t *testing.T) {
		data := ReadData("../example.dat")
		// fmt.Println(data)
		wh := WarehouseFromStr(data)
		fmt.Println(wh)
		if _, exists := wh.walls[NewLocation(0, 3)]; !exists  {
			t.Errorf("Expected a piece of wall")
		} 
		if _, exists := wh.walls[NewLocation(2, 5)]; !exists  {
			t.Errorf("Expected a piece of wall")
		} 
		if _, exists := wh.walls[NewLocation(wh.dimensions.x-1, 3)]; !exists  {
			t.Errorf("Expected a piece of wall")
		} 
		if _, exists := wh.boxes[NewLocation(3, 1)]; !exists  {
			t.Errorf("Expected a box")
		} 
		if _, exists := wh.boxes[NewLocation(5, 8)]; !exists  {
			t.Errorf("Expected a box")
		} 

		if len(wh.robotPath.pointers) != 700 {
			t.Errorf("Path is wrong")
		}
	})

}
