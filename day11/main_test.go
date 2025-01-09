package main

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func Test_Map(t *testing.T) {
	t.Run("DoNGenMap", func(t *testing.T) {
		// firstGen := []string{"125", "17"}

		stones := []string{"125", "17"}
		stonesMap := dedup(&stones)
		ngen := 25
		sum := int64(1)
		sum = stonesMap.DoNgenMap(ngen)
		fmt.Printf("Stone %v after %d generations. Number of stones: %d\n", stones, ngen, sum)
	})

	t.Run("NextGenMap", func(t *testing.T) {
		input := strings.Split("8096 1 8096 16192 2 0 2 4 8096 1 8096 16192 16192 1 18216 12144", " ")
		expected := strings.Split("80 96 2024 80 96 32772608 4048 1 4048 8096 80 96 2024 80 96 32772608 32772608 2024 36869184 24579456", " ")
		expectedMap := dedup(&expected)
		if expectedMap.CountStones() != 20 {
			t.Errorf("CountSum should be 20, is %d", expectedMap.CountStones())
		}

		gen := dedup(&input)
		nextGen := (&gen).NextGenMap()
		if nextGen.CountStones() != 20 {
			t.Errorf("CountSum should be 20, is %d", nextGen.CountStones())
		}
		fmt.Println(nextGen)
		if !reflect.DeepEqual(nextGen, expectedMap) {
			t.Errorf("Result not as expected")
		}

		input = strings.Split("512 72 2024 2 0 2 4 2867 6032", " ")
		expected = strings.Split("1036288 7 2 20 24 4048 1 4048 8096 28 67 60 32", " ")
		expectedMap = dedup(&expected)

		gen = dedup(&input)
		nextGen = (&gen).NextGenMap()
		fmt.Println(nextGen)
		if !reflect.DeepEqual(nextGen, expectedMap) {
			t.Errorf("Result not as expected")
		}

		if gen.CountStones() != 9 {
			t.Errorf("CountSum should be 9, but is %d", gen.CountStones())
		}

		if nextGen.CountStones() != 13 {
			t.Errorf("CountSum should be 13, but is %d", nextGen.CountStones())
		}

		input = strings.Split("512072 1 20 24 28676032", " ")
		expected = strings.Split("512 72 2024 2 0 2 4 2867 6032", " ")
		expectedMap = dedup(&expected)

		gen = dedup(&input)
		nextGen = (&gen).NextGenMap()
		fmt.Println(nextGen)
		if !reflect.DeepEqual(nextGen, expectedMap) {
			t.Errorf("Result not as expected 2")
		}

	})

	t.Run("applyRules", func(t *testing.T) {
		if applyRules("0")[0] != "1" {
			t.Errorf("0 Rule failed")
		}
		if applyRules("2")[0] != "4048" {
			t.Errorf("times 2024 Rule failed")
		}
		gen := applyRules("2234")
		if !reflect.DeepEqual(gen, []string{"22", "34"}) {
			t.Errorf("split Rule failed")
		}
	})

}

func Test_Walk(t *testing.T) {
	t.Run("DoNGen", func(t *testing.T) {
		// firstGen := []string{"125", "17"}

		stones := []string{"125", "17"}
		ngen := 25
		sum := int64(0)
		sum = DoNgen(stones, ngen)
		fmt.Printf("Stone %v after %d generations. Number of stones: %d\n", stones, ngen, sum)
        if sum != 55312 {
            t.Errorf("Expected 55312 stones, got %d", sum)
        }
	})

	t.Run("Walk", func(t *testing.T) {
		// firstGen := []string{"125", "17"}

		var genCound int
		var stone string

		stone = "1"
		genCound = 25
		sum := int64(1)
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
