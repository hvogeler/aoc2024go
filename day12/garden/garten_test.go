package garden

import (
	"fmt"
	"testing"
)

const example3 = `RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`

func Test_FromStr(t *testing.T) {
	t.Run("DoNGenMap", func(t *testing.T) {
		garden := GardenFromStr(example3)
		if garden.dimensions.rows != 10 {
			t.Errorf("Expected 10 rows, got %d", garden.dimensions.rows)
		}
		if garden.dimensions.cols != 10 {
			t.Errorf("Expected 10 columns, got %d", garden.dimensions.cols)
		}
		if !garden.area[2][2].Equals(Plot{PlantType('R'), Location{2, 2}, []Fence{}})  {
			t.Errorf("Expected %v got %v", Plot{PlantType('R'), Location{2, 2}, []Fence{}}, garden.area[2][2])
		}
		fmt.Println(garden)
	})

}