package warehouse

import (
	"fmt"
	"testing"
)

func Test_xmas_tree(t *testing.T) {
	t.Run("Example1 Data", func(t *testing.T) {
		data := ReadData("../example1.dat")
		// fmt.Println(data)
		wh := WarehouseFromStr(data)
		for i := 0; i < wh.robotPath.Length(); i++ {
			ptr := wh.robotPath.NextPointer()
			wh.Move(&wh.robot.position, ptr)
			fmt.Println(ptr.String())
			fmt.Println(wh)
		}
		fmt.Printf("Sum of box coords: %d\n", wh.SumBoxCoords())
		if wh.SumBoxCoords() != 2028 {
			t.Errorf("Expected Sum 2028, got %d", wh.SumBoxCoords())
		}
	})

	t.Run("Example Data", func(t *testing.T) {
		data := ReadData("../example.dat")
		// fmt.Println(data)
		wh := WarehouseFromStr(data)
		wh.GoRobotGo()
		fmt.Printf("Sum of box coords: %d\n", wh.SumBoxCoords())
		if wh.SumBoxCoords() != 10092 {
			t.Errorf("Expected Sum 2028, got %d", wh.SumBoxCoords())
		}

	})

	t.Run("Warehouse", func(t *testing.T) {
		data := ReadData("../example.dat")
		// fmt.Println(data)
		wh := WarehouseFromStr(data)
		fmt.Println(wh)
		if _, exists := wh.walls[NewLocation(0, 3)]; !exists {
			t.Errorf("Expected a piece of wall")
		}
		if _, exists := wh.walls[NewLocation(2, 5)]; !exists {
			t.Errorf("Expected a piece of wall")
		}
		if _, exists := wh.walls[NewLocation(wh.dimensions.x-1, 3)]; !exists {
			t.Errorf("Expected a piece of wall")
		}
		if _, exists := wh.boxes[NewLocation(3, 1)]; !exists {
			t.Errorf("Expected a box")
		}
		if _, exists := wh.boxes[NewLocation(5, 8)]; !exists {
			t.Errorf("Expected a box")
		}

		if len(wh.robotPath.pointers) != 700 {
			t.Errorf("Path is wrong")
		}
	})

}
