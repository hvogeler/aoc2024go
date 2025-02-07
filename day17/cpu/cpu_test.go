package cpu

import (
	"fmt"
	"testing"
)

func Test_1(t *testing.T) {
	t.Run("Example Data1", func(t *testing.T) {
		data := ReadData("../example1.dat")
		fmt.Println(data)
		cpu := InitialProgramLoad(data)
		fmt.Println(cpu.DisAssemble(-1))
	})
}
