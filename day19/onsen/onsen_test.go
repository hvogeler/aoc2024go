package onsen

import (
	"fmt"
	"testing"
)

func Test_1(t *testing.T) {
	t.Run("Example Data1", func(t *testing.T) {
		data := ReadData("../example1.dat")
		fmt.Println(data)
		onsen := OnsenFromStr(data)
		
		if len(onsen.designs) != 8 {
			t.Errorf("Number of designs wrong")
		}
		
		if len(onsen.pattern) != 8 {
			t.Errorf("Number of patterns wrong")
		}

		if !onsen.ContainsPattern("bwu") {
			t.Errorf("Pattern bwu missing")
		}

		if !onsen.ContainsPattern("br") {
			t.Errorf("Pattern br missing")
		}

		if !onsen.ContainsPattern("r") {
			t.Errorf("Pattern r missing")
		}
	})
}
