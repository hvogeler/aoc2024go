package maze

import (
	"fmt"
	"testing"
)

func Test_FromString(t *testing.T) {
	t.Run("Example Data1", func(t *testing.T) {
		data := ReadData("../example1.dat")
		fmt.Println(data)
		maze := MazeFromStr(data)
		fmt.Println(maze)
	})
}

func Test_Move(t *testing.T) {
	t.Run("Move1", func(t *testing.T) {
		data := ReadData("../example1.dat")
		// fmt.Println(data)
		maze := MazeFromStr(data)
		fmt.Println(maze)
		for i := 0; maze.CountAlive() > 0; i++ {
			maze.MoveReindeer()
			fmt.Println(maze)
			fmt.Println()
		}
		if maze.lowScore != 7036 {
			t.Errorf("Expected low score 7036, got %d", maze.lowScore)
		}
	})

	t.Run("Move2", func(t *testing.T) {
		data := ReadData("../example2.dat")
		// fmt.Println(data)
		maze := MazeFromStr(data)
		fmt.Println(maze)
		for i := 0; maze.CountAlive() > 0; i++ {
			maze.MoveReindeer()
			fmt.Println(maze)
		}
		if maze.lowScore != 11048 {
			t.Errorf("Expected low score 11048, got %d", maze.lowScore)
		}
	})

	t.Run("Move3", func(t *testing.T) {
		data := ReadData("../example3.dat")
		// fmt.Println(data)
		maze := MazeFromStr(data)
		fmt.Println(maze)
		for i := 0; maze.CountAlive() > 0; i++ {
			maze.MoveReindeer()
			fmt.Println(maze)
		}
		// part2: 149
		if maze.lowScore != 21148 {
			t.Errorf("Expected low score 21148, got %d", maze.lowScore)
		}
	})

	t.Run("Move4", func(t *testing.T) {
		data := ReadData("../example4.dat")
		// fmt.Println(data)
		maze := MazeFromStr(data)
		fmt.Println(maze)
		for i := 0; maze.CountAlive() > 0; i++ {
			maze.MoveReindeer()
			fmt.Println(maze)
		}

		if maze.lowScore != 4013 {
			t.Errorf("Expected low score 4013, got %d", maze.lowScore)
		}
	})

	t.Run("Move5", func(t *testing.T) {
		data := ReadData("../example5.dat")
		// fmt.Println(data)
		maze := MazeFromStr(data)
		fmt.Println(maze)
		for i := 0; maze.CountAlive() > 0; i++ {
			maze.MoveReindeer()
			fmt.Println(maze)
		}

		// part2: 413
		if maze.lowScore != 5078 {
			t.Errorf("Expected low score 4013, got %d", maze.lowScore)
		}
	})

	t.Run("Move6", func(t *testing.T) {
		data := ReadData("../example6.dat")
		// fmt.Println(data)
		maze := MazeFromStr(data)
		fmt.Println(maze)
		for i := 0; maze.CountAlive() > 0; i++ {
			maze.MoveReindeer()
		}
		fmt.Println(maze)

		// part2: 264
		if maze.lowScore != 21110 {
			t.Errorf("Expected low score 21110, got %d", maze.lowScore)
		}
	})

	t.Run("Move7", func(t *testing.T) {
		data := ReadData("../example7.dat")
		// fmt.Println(data)
		maze := MazeFromStr(data)
		fmt.Println(maze)
		for i := 0; maze.CountAlive() > 0; i++ {
			maze.MoveReindeer()
			// fmt.Println(maze)
		}
		fmt.Println(maze)

		// part2: 514
		if maze.lowScore != 41210 {
			t.Errorf("Expected low score 41210, got %d", maze.lowScore)
		}
	})

	t.Run("Move8", func(t *testing.T) {
		data := ReadData("../example8.dat")
		// fmt.Println(data)
		maze := MazeFromStr(data)
		fmt.Println(maze)
		for i := 0; maze.CountAlive() > 0; i++ {
			maze.MoveReindeer()
			// fmt.Println(maze)
		}
		fmt.Println(maze)
		fmt.Println(maze.PrintCheapestTrack())
		// part2: 514
		if maze.lowScore != 4021 {
			t.Errorf("Expected low score 4021, got %d", maze.lowScore)
		}
	})

	t.Run("testdata", func(t *testing.T) {
		data := ReadData("../testdata.dat")
		// fmt.Println(data)
		maze := MazeFromStr(data)
		fmt.Println(maze)
		i := 0
		for i = 0; maze.CountAlive() > 0; i++ {
			maze.MoveReindeer()
			if i%100 == 0 {
				fmt.Println(maze)
				fmt.Println()
			}
		}
		fmt.Printf("Reindeer cloned: %d\n", len(maze.reindeers))
		fmt.Println(maze.PrintCheapestTrack())


		// 153516 is too high
		if maze.lowScore != 11048 {
			t.Errorf("Expected low score 11048, got %d", maze.lowScore)
		}
	})
}
