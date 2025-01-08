package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	bytes, err := os.ReadFile("testdata.dat")
	if err != nil {
		panic(err)
	}
	raw := string(bytes)
	var data string = string(raw)
	firstGen := strings.Split(data, " ")

	n := DoNgen(firstGen, 25)
	fmt.Println("Stones: ", n)
	// Part 1: 199946

	sum := 0
	for i, stone := range firstGen {
		fmt.Printf("Stone %s, %d of %d\n", stone, i + 1, len(firstGen))
		n = DoNgen([]string{stone}, 75)
		fmt.Printf("Stone %s, %d of %d, Stones %d", stone, i + i, len(firstGen), n)
		sum += n
	}
	fmt.Println("Stones: ", sum)


}

func Walk(stone string, agg *int, currentGen int, maxGen int) {
	fmt.Printf("Walk Stone '%s' at Generation %d. Agg = %d\n", stone, currentGen, *agg)
	stones := []string{stone}

	i := 0
	for ; len(stones) == 1; i++ {
		stones = NextGen(stones)
	}

	*agg += 1

	if currentGen > maxGen - 2 {
		return
	}
	
	Walk(stones[0], agg, currentGen + i, maxGen)
	Walk(stones[1], agg, currentGen + i, maxGen)
}

func NextGen(currentGen []string) []string {
	nextGen := []string{}
	for _, inscription := range currentGen {
		if inscription == "0" {
			nextGen = append(nextGen, "1")
			continue
		}

		if len(inscription)%2 == 0 {
			runes := []rune(inscription)
			leftHalf := runes[:len(inscription)/2]
			rightHalf := runes[len(inscription)/2:]
			nextGen = append(nextGen, ReduceZeros(string(leftHalf)))
			nextGen = append(nextGen, ReduceZeros(string(rightHalf)))
			continue
		}

		v, err := strconv.Atoi(inscription)
		if err == nil {
			v *= 2024
			nextGen = append(nextGen, ReduceZeros(fmt.Sprintf("%d", v)))
		}
	}
	return nextGen
}

func ReduceZeros(s string) string {
	v, err := strconv.Atoi(s)
	if err == nil {
		return fmt.Sprintf("%d", v)
	}
	return s
}

func DoNgen(firstGen []string, n int) int {
	currentGen := firstGen
	for i := 0; i < n; i++ {
		nextGen := NextGen(currentGen)
		fmt.Printf("%d Gen %d stones", i + 1, len(nextGen))
		if len(nextGen) < 20 {
			fmt.Printf(" %v\n", nextGen)
		} else {
			fmt.Println()
		}
		currentGen = nextGen
	}

	return len(currentGen)
}