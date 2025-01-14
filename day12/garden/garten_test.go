package garden

import (
	"fmt"
	"reflect"
	"testing"
)

const example1 = `RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`

const example2 = `OOOOO
OXOXO
OOOOO
OXOXO
OOOOO`

func Test_corners(t *testing.T) {
	t.Run("plot no neighbor", func(t *testing.T) {
		var plot Plot

		plot.neighbors = [4]*Plot{}
		if len(plot.Corners()) != 4 {
			t.Errorf("Expected 4 corners, got %d", len(plot.Corners()))
		}
		if plot.Corners()[0].cornerType != convex || plot.Corners()[0].orientation != topLeft {
			t.Errorf("Wrong corned data")
		}
	})

	t.Run("plot all concav (+)", func(t *testing.T) {
		var plot Plot
		var nabove Plot

		plot.neighbors = [4]*Plot{}
		plot.neighbors[above] = &nabove
		plot.neighbors[below] = &nabove
		plot.neighbors[right] = &nabove
		plot.neighbors[left] = &nabove
		if len(plot.Corners()) != 4 {
			t.Errorf("Expected 4 corners, got %d", len(plot.Corners()))
		}
        for _, corner := range plot.Corners() {
            if corner.cornerType != concave {
                t.Errorf("Expected concave corner")
            }
        }
	})

	t.Run("plot all neighbors", func(t *testing.T) {
		var plot Plot
        var n Plot
        var n1 Plot
        for i := 0; i < 4; i++ {
            n.neighbors[i] = &n1
        }
		plot.neighbors = [4]*Plot{}
        for i := 0; i < 4; i++ {
            plot.neighbors[i] = &n
        }

		if len(plot.Corners()) != 0 {
			t.Errorf("Expected 0 corners, got %d", len(plot.Corners()))
		}
	})
}

func Test_FenceCost2(t *testing.T) {
	t.Run("FenceCost1", func(t *testing.T) {
		garden := GardenFromStr(example1)
		garden.findRegions()
        fmt.Println(garden.regions[0])
		if len(garden.regions[0].corners) != 10 {
			t.Errorf("Expected 10 corners, got %d", len(garden.regions[0].corners))
		}
        fmt.Println(garden.regions[1])
		if len(garden.regions[1].corners) != 4 {
			t.Errorf("Expected 4 corners, got %d", len(garden.regions[1].corners))
		}
        fmt.Println(garden.regions[2])
		if len(garden.regions[2].corners) != 22 {
			t.Errorf("Expected 22 corners, got %d", len(garden.regions[1].corners))
		}
        fmt.Println(garden.regions[3])
		if len(garden.regions[3].corners) != 12 {
			t.Errorf("Expected 12 corners, got %d", len(garden.regions[1].corners))
		}
        fmt.Println(garden.regions[4])
		if len(garden.regions[4].corners) != 10 {
			t.Errorf("Expected 10 corners, got %d", len(garden.regions[1].corners))
		}
        fmt.Println(garden.regions[6])
		if len(garden.regions[6].corners) != 4 {
			t.Errorf("Expected 4 corners, got %d", len(garden.regions[1].corners))
		}
        fmt.Println(garden.regions[8])
		if len(garden.regions[8].corners) != 16 {
			t.Errorf("Expected 16 corners, got %d", len(garden.regions[1].corners))
		}
        fmt.Println(garden.regions[10])
		if len(garden.regions[10].corners) != 6 {
			t.Errorf("Expected 6 corners, got %d", len(garden.regions[1].corners))
		}
	})
}

func Test_FenceCost(t *testing.T) {
	t.Run("FenceCost1", func(t *testing.T) {
		garden := GardenFromStr(example1)
		garden.findRegions()
		fenceCost := garden.regions[0].FenceCost()
		if fenceCost != 216 {
			t.Errorf("Expected fencecost 216, got %d", fenceCost)
		}
	})

	t.Run("FenceCostO", func(t *testing.T) {
		garden := GardenFromStr(example2)
		garden.findRegions()
		fenceCost := garden.regions[0].FenceCost()
		if fenceCost != 756 {
			t.Errorf("Expected fencecost 756, got %d", fenceCost)
		}
	})

	t.Run("FenceCostGarden1", func(t *testing.T) {
		garden := GardenFromStr(example1)
		garden.findRegions()
		fenceCost := garden.FenceCost()
		if fenceCost != 1930 {
			t.Errorf("Expected fencecost 1930, got %d", fenceCost)
		}
	})

	t.Run("FenceCostGarden2", func(t *testing.T) {
		garden := GardenFromStr(example2)
		garden.findRegions()
		fenceCost := garden.FenceCost()
		if fenceCost != 772 {
			t.Errorf("Expected fencecost 772, got %d", fenceCost)
		}
	})
}

func Test_OOO(t *testing.T) {
	t.Run("findRegions", func(t *testing.T) {
		garden := GardenFromStr(example2)
		garden.findRegions()
		if len(garden.regions) != 5 {
			t.Errorf("Expected 5 regions, got %d", len(garden.regions))
		}
		for _, region := range garden.regions {
			fmt.Println(region)
		}
	})
}

func Test_FromStr(t *testing.T) {
	t.Run("findRegions", func(t *testing.T) {
		garden := GardenFromStr(example1)
		garden.findRegions()
		if len(garden.regions) != 11 {
			t.Errorf("Expected 11 regions, got %d", len(garden.regions))
		}
		for _, region := range garden.regions {
			fmt.Println(region)
		}

	})

	t.Run("WalkPlots", func(t *testing.T) {
		garden := GardenFromStr(example1)
		region1 := new(Region)
		garden.area[0][0].WalkPlot(region1)
		if len(region1.plots) != 12 {
			t.Errorf("Expected region 'R' to contain 12 plots, but got %d", len(region1.plots))
		}

		region2 := new(Region)
		garden.area[2][3].WalkPlot(region2)
		if len(region2.plots) != 12 {
			t.Errorf("Expected region 'R' to contain 12 plots, but got %d", len(region2.plots))
		}
		region1.Sort()
		region2.Sort()
		if !reflect.DeepEqual(region1, region2) {
			t.Errorf("Regions should be equal")
		}

		fmt.Println(region1.String())

		region3 := new(Region)
		garden.area[1][6].WalkPlot(region3)
		fmt.Println(region3.String())
	})

	t.Run("FromStr", func(t *testing.T) {
		garden := GardenFromStr(example1)
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
