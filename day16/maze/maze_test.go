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
	
	t.Run("testdata", func(t *testing.T) {
		data := ReadData("../testdata.dat")
		// fmt.Println(data)
		maze := MazeFromStr(data)
		fmt.Println(maze)
		i := 0
		for i = 0; maze.CountAlive() > 0; i++ {
			maze.MoveReindeer()
			fmt.Println(maze)
		}
        fmt.Printf("Reindeer cloned: %d\n", len(maze.reindeers))
        // 153516 is too high
		if maze.lowScore != 11048 {
			t.Errorf("Expected low score 11048, got %d", maze.lowScore)
		}
	})
}
