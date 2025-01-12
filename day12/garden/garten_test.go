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
	t.Run("WalkPlots", func(t *testing.T) {
		garden := GardenFromStr(example3)
		region := new(Region)
		garden.area[0][0].WalkPlot(region)
		if len(*region) != 12 {
			t.Errorf("Expected region 'R' to contain 12 plots, but got %d", len(*region))
		}
		garden.area[2][3].WalkPlot(region)
		if len(*region) != 12 {
			t.Errorf("Expected region 'R' to contain 12 plots, but got %d", len(*region))
		}
		// garden.area[0][0].WalkPlot(region)
		// if len(*region) != 12 {
		// 	t.Errorf("Expected region 'R' to contain 12 plots, but got %d", len(*region))
		// }
	})

	t.Run("FromStr", func(t *testing.T) {
		garden := GardenFromStr(example3)
		if garden.dimensions.rows != 10 {
			t.Errorf("Expected 10 rows, got %d", garden.dimensions.rows)
		}
		if garden.dimensions.cols != 10 {
			t.Errorf("Expected 10 columns, got %d", garden.dimensions.cols)
		}
		plot := garden.area[2][2]
		if plot.plantType != PlantType('R') {
			t.Errorf("Expected planttype 'R', got %s", plot.plantType)
		}
		if plot.neighbors[above] == nil || plot.neighbors[above].plantType != PlantType('R') {
			t.Errorf("Expected neighbor above of planttype 'R'")
		}
		if plot.neighbors != [4]*Plot{&garden.area[1][2], &garden.area[2][3], &garden.area[3][2], nil} {
			t.Errorf("Expected other neighbors")
		}

		plot = garden.area[9][0]
		if plot.plantType != PlantType('M') {
			t.Errorf("Expected planttype 'M', got %s", plot.plantType)
		}
		if plot.neighbors != [4]*Plot{&garden.area[8][0], &garden.area[9][1], nil, nil} {
			t.Errorf("Expected other neighbors")
		}

		plot = garden.area[4][9]
		if plot.plantType != PlantType('E') {
			t.Errorf("Expected planttype 'E', got %s", plot.plantType)
		}
		if plot.neighbors != [4]*Plot{nil, nil, &garden.area[5][9], nil} {
			t.Errorf("Expected other neighbors")
		}

		plot = garden.area[0][1]
		if plot.plantType != PlantType('R') {
			t.Errorf("Expected planttype 'R', got %s", plot.plantType)
		}
		if plot.neighbors[right].location != (Location{0, 2}) {
			t.Errorf("Expected neighbor location (0, 2), got %s", plot.neighbors[right].location)
		}

		fmt.Println(garden)
	})

}
