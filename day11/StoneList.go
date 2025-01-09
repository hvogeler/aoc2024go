package main

import (
	"fmt"
	"strconv"
)


func DoNgen(firstGen []string, n int) int64 {
	currentGen := firstGen
	for i := 0; i < n; i++ {
		nextGen := NextGen(currentGen)
		// fmt.Printf("%d Gen %d stones, stones %d", i+1, len(nextGen), len(nextGen))
		fmt.Printf("Gen %d: %d stones  ", i+1, len(nextGen))
		// fmt.Printf("%d\n", len(nextGen))
		if len(nextGen) < 150 {
			fmt.Printf(" %v\n", nextGen)
		} else {
			fmt.Println()
		}
		currentGen = nextGen
	}

	return int64(len(currentGen))
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
