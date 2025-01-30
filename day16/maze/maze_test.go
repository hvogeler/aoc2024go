package maze

import (
	"fmt"
	"testing"
)

func Test_FromString2(t *testing.T) {
	t.Run("Example Data1", func(t *testing.T) {
		data := ReadData("../example1.dat")
		fmt.Println(data)
		maze := MazeFromStr(data)
		fmt.Println(maze)
	})
}
