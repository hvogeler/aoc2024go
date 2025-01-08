package main

import (
	"fmt"
	"strings"
	"testing"
)

func Test_Walk(t *testing.T) {

	t.Run("Walk", func(t *testing.T) {
		// firstGen := []string{"125", "17"}

		var genCound int
		var stone string

		stone = "125"
		genCound = 50
		sum := int64(1)
		// if len(stone) % 2 == 0 {
		// 	sum = 1
		// }
		Walk(stone, &sum, 0, genCound)
		fmt.Printf("Stone %s after %d generations. Number of stones: %d\n", stone, genCound, sum)
	})

}

func Test_play(t *testing.T) {

	t.Run("nextGen", func(t *testing.T) {
		// firstGen := []string{"125", "17"}

		var genCound int
		var stone string

		stone = "17"
		genCound = 6
		n := DoNgen([]string{stone}, genCound)
		fmt.Printf("Number of stones for %s after %d generations: %d\n", stone, genCound, n)
	})

}

func Test_nextgen(t *testing.T) {

	t.Run("nextGen", func(t *testing.T) {
		firstGen := []string{"125", "17"}

		secondGen := NextGen(firstGen)
		fmt.Println(secondGen)
		if strings.Join(secondGen, " ") != "253000 1 7" {
			t.Errorf("Second Gen Wrong")
		}

		thirdGen := NextGen(secondGen)
		fmt.Println(thirdGen)
		if strings.Join(thirdGen, " ") != "253 0 2024 14168" {
			t.Errorf("Third Gen Wrong: %v", thirdGen)
		}

		fourthGen := NextGen(thirdGen)
		fmt.Println(fourthGen)
		if strings.Join(fourthGen, " ") != "512072 1 20 24 28676032" {
			t.Errorf("fourthGenWrong: %v", fourthGen)
		}

		fifthGen := NextGen(fourthGen)
		fmt.Println(fifthGen)
		if strings.Join(fifthGen, " ") != "512 72 2024 2 0 2 4 2867 6032" {
			t.Errorf("fifthGen Wrong: %v", fifthGen)
		}

		sixthGen := NextGen(fifthGen)
		fmt.Println(sixthGen)
		if strings.Join(sixthGen, " ") != "1036288 7 2 20 24 4048 1 4048 8096 28 67 60 32" {
			t.Errorf("fifthGen Wrong: %v", sixthGen)
		}

		seventhGen := NextGen(sixthGen)
		fmt.Println(seventhGen)
		if strings.Join(seventhGen, " ") != "2097446912 14168 4048 2 0 2 4 40 48 2024 40 48 80 96 2 8 6 7 6 0 3 2" {
			t.Errorf("fifthGen Wrong: %v", seventhGen)
		}
	})

	t.Run("25Gen", func(t *testing.T) {
		firstGen := []string{"125", "17"}
		nStones := DoNgen(firstGen, 6)
		if nStones != 22 {
			t.Errorf("Expected 22 stones, got %d", nStones)
		}

		nStones = DoNgen(firstGen, 25)
		if nStones != 55312 {
			t.Errorf("Expected 55312 stones, got %d", nStones)
		}

	})
}
