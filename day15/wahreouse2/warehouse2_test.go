package warehouse2

import (
	wh1 "day15/warehouse"
	_ "embed"
	"fmt"
	"testing"
)

//go:embed example2.dat
var example2 string

func Test_FromString2(t *testing.T) {
	t.Run("Example Data", func(t *testing.T) {
		data := wh1.ReadData("../example.dat")
		// fmt.Println(data)
		wh := WarehouseFromStr(data)
		fmt.Println(wh)

		item, exists := wh.ItemAt(7, 0)
		if exists {
			if item.Item() != WallItem {
				t.Errorf("Expected Wall")
			}
		}

		item, exists = wh.ItemAt(7, 1)
		if exists {
			if item.Item() != BoxItem || item.PositionRight() != NewLocation(7, 1) {
				t.Errorf("Expected Box")
			}
		} else {
			t.Errorf("item does not exist")
		}

		item, exists = wh.ItemAt(6, 1)
		if exists && item.Item() != BoxItem || item.PositionLeft() != NewLocation(6, 1) {
			t.Error("Expected Box")
		}
	})

	t.Run("Example2 Data", func(t *testing.T) {
		// data := wh1.ReadData("../example.dat")
		// fmt.Println(data)
		wh := WarehouseFromStr(example2)
		fmt.Println(wh)
		for i := 0; i < wh.robotPath.Length(); i++ {
			ptr := wh.robotPath.NextPointer()
			wh.Move(&wh.robot.position, ptr, 0)
			fmt.Println(wh)
		}
		if wh.SumBoxCoords() != 105+207+306 {
			t.Errorf("Espected Sum 618, got %d", wh.SumBoxCoords())
		}
	})

}

func Test_DroppingBoxes(t *testing.T) {

	t.Run("Example Data 3", func(t *testing.T) {
		data := wh1.ReadData("../example3.dat")
		wh := WarehouseFromStr(data)
		fmt.Println(wh)
		for i := 0; i < wh.robotPath.Length(); i++ {
			ptr := wh.robotPath.NextPointer()
            fmt.Printf("\n%d. %s\n", i, ptr)
			wh.Move(&wh.robot.position, ptr, 0)
			fmt.Println(wh)
		}
	})
}

func Test_FromString3(t *testing.T) {

	t.Run("Example Data", func(t *testing.T) {
		data := wh1.ReadData("../testdata.dat")
		wh := WarehouseFromStr(data)
		fmt.Println(wh)
		for i := 0; i < wh.robotPath.Length(); i++ {
			ptr := wh.robotPath.NextPointer()
            fmt.Printf("\n%d. %s\n", i, ptr)
			wh.Move(&wh.robot.position, ptr, 0)
			fmt.Println(wh)
		}
		if wh.SumBoxCoords() != 9021 {
			t.Errorf("Espected Sum 9021, got %d", wh.SumBoxCoords())
		}
	})

	t.Run("Example Data gogo", func(t *testing.T) {
		data := wh1.ReadData("../example.dat")
		wh := WarehouseFromStr(data)
		fmt.Println(wh)
        wh.GoRobotGo()
		if wh.SumBoxCoords() != 9021 {
			t.Errorf("Espected Sum 9021, got %d", wh.SumBoxCoords())
		}
	})
}
