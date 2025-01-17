package main

import (
	"fmt"
	"os"
	"testing"
)

func Test_t1(t *testing.T) {
	t.Run("read good1", func(t *testing.T) {
		bytes, err := os.ReadFile("example.dat")
		if err != nil {
			panic(err)
		}
		data := string(bytes)
		games := games(data)
		fmt.Println(games[0])
		if len(games) != 4 {
			t.Errorf("expected 4 games, got %d", len(games))
		}
		
	})

}
