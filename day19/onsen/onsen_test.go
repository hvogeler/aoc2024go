package onsen

import (
	"fmt"
	"testing"
)

func Test_long_patterns(t *testing.T) {
	t.Run("Check Designs largest sub wrong", func(t *testing.T) {
		data := ReadData("../example2.dat")
		fmt.Println(data)
		onsen := OnsenFromStr(data)
		d := onsen.CheckDesign(onsen.designs[0])
		fmt.Println(d)
		if !d.isPossible {
			t.Errorf("Is Possible: bw, up, bw")
		}
	})
}

func Test_all_patterns(t *testing.T) {
	t.Run("Check Designs1", func(t *testing.T) {
		data := ReadData("../example1.dat")
		fmt.Println(data)
		onsen := OnsenFromStr(data)
		onsen.CheckDesigns()
		all, possible, impossible := onsen.CountDesigns()

		if all != 8 {
			t.Errorf("Count All wrong")
		}
		if possible != 6 {
			t.Errorf("Count possible wrong")
		}
		if impossible != 2 {
			t.Errorf("Count impossible wrong")
		}
	})
}

func Test_1(t *testing.T) {
	t.Run("Check Designs1", func(t *testing.T) {
		data := ReadData("../example1.dat")
		fmt.Println(data)
		onsen := OnsenFromStr(data)
		d := onsen.CheckDesign(onsen.designs[0])
		fmt.Println(d)
		if len(d.designPattern) != 3 {
			t.Errorf("Designpatterns found wrong")
		}
		if d.designPattern[0] != "br" || d.designPattern[2] != "r" {
			t.Errorf("Designpatterns found wrong")
		}
		if !d.isPossible {
			t.Errorf("Designpatterns should be possible")
		}
	})

	t.Run("Check Designs2", func(t *testing.T) {
		data := ReadData("../example1.dat")
		fmt.Println(data)
		var onsen Onsen
		var d *Design

		onsen = OnsenFromStr(data)
		d = onsen.CheckDesign(onsen.designs[3])
		fmt.Println(d)
		if len(d.designPattern) != 4 {
			t.Errorf("Designpatterns found wrong")
		}
		if d.designPattern[0] != "r" || d.designPattern[2] != "gb" || d.designPattern[3] != "r" {
			t.Errorf("Designpatterns found wrong")
		}
		if !d.isPossible {
			t.Errorf("Designpatterns should be possible")
		}
	})

	t.Run("Check Designs3", func(t *testing.T) {
		data := ReadData("../example1.dat")
		fmt.Println(data)
		var onsen Onsen
		var d *Design

		onsen = OnsenFromStr(data)
		d = onsen.CheckDesign(onsen.designs[1])
		fmt.Println(d)
		if len(d.designPattern) != 4 {
			t.Errorf("Designpatterns found wrong")
		}
		if d.designPattern[0] != "b" || d.designPattern[2] != "g" || d.designPattern[3] != "r" {
			t.Errorf("Designpatterns found wrong")
		}
		if !d.isPossible {
			t.Errorf("Designpatterns should be possible")
		}
	})

	t.Run("Check Designs4", func(t *testing.T) {
		data := ReadData("../example1.dat")
		fmt.Println(data)
		var onsen Onsen
		var d *Design

		onsen = OnsenFromStr(data)
		d = onsen.CheckDesign(onsen.designs[5])
		fmt.Println(d)
		if len(d.designPattern) != 4 {
			t.Errorf("Designpatterns found wrong")
		}
		if d.designPattern[0] != "bwu" || d.designPattern[2] != "r" || d.designPattern[3] != "g" {
			t.Errorf("Designpatterns found wrong")
		}
		if !d.isPossible {
			t.Errorf("Designpatterns should be possible")
		}
	})

	t.Run("Check Designs5", func(t *testing.T) {
		data := ReadData("../example1.dat")
		fmt.Println(data)
		var onsen Onsen
		var d *Design

		onsen = OnsenFromStr(data)
		d = onsen.CheckDesign(onsen.designs[7])
		fmt.Println(d)
		if d.isPossible {
			t.Errorf("Designpatterns should NOT be possible")
		}
	})
}

func Test_2(t *testing.T) {
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

		if onsen.LongestPatternLength() != 3 {
			t.Errorf("Longest pattern length should be 3, got %d", onsen.LongestPatternLength())
		}

	})
}
