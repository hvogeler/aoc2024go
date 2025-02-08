package memspace


import (
	"fmt"
	"testing"
)

func Test_1(t *testing.T) {
	t.Run("Example Data1", func(t *testing.T) {
		data := ReadData("../example1.dat")
		fmt.Println(data)
		memSpace := MemSpaceFromStr(data, 7, 7, 12)
		fmt.Println(memSpace)
	})
}
