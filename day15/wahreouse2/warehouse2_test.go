package warehouse2

import (
	"fmt"
	"testing"
	wh1 "day15/warehouse"
)

func Test_FromString2(t *testing.T) {
t.Run("Example Data", func(t *testing.T) {
	data := wh1.ReadData("../example.dat")
	// fmt.Println(data)
	wh := WarehouseFromStr(data)
	fmt.Println(wh)
})
}

