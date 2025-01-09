package main

import (
	"fmt"
	"strconv"
)


type StoneMap map[string]int

func (stones StoneMap) CountStones() int64 {
	sum := int64(0)
	for _, count := range stones {
		sum += int64(count)
	}
	return sum
}

func dedup(stones *[]string) StoneMap {
	m := make(map[string]int)
	for _, stone := range *stones {
		m[stone] = m[stone] + 1
	}
	return m
}

func (firstGen StoneMap) DoNgenMap(generations int) int64 {
	currentGen := firstGen
	for i := 0; i < generations; i++ {
		nextGen := (&currentGen).NextGenMap()
		fmt.Printf("Gen %d: %d stones, MapLen: %d ", i+1, nextGen.CountStones(), len(nextGen))
		if len(nextGen) < 15 {
			fmt.Printf(" %v\n", nextGen)
		} else {
			fmt.Println()
		}
		currentGen = nextGen
	}

	return int64(currentGen.CountStones())
}

func (currentGen *StoneMap) NextGenMap() StoneMap {
	nextGen := make(StoneMap)
	for inscription, count := range *currentGen {
		newInscription := applyRules(inscription)
		switch len(newInscription) {
		case 1:
			nextGen[newInscription[0]] += count
		case 2:
			nextGen[newInscription[0]] += count
			nextGen[newInscription[1]] += count
		default:
			panic("Rule result must be len 2 or less")
		}
	}

	return nextGen
}

func applyRules(inscription string) []string {
	newInscription := []string{}
	if inscription == "0" {
		newInscription = []string{"1"}
		return newInscription
	}

	if len(inscription)%2 == 0 {
		runes := []rune(inscription)
		leftHalf := runes[:len(inscription)/2]
		rightHalf := runes[len(inscription)/2:]
		newInscription = []string{ReduceZeros(string(leftHalf)), ReduceZeros(string(rightHalf))}
		return newInscription
	}

	v, err := strconv.Atoi(inscription)
	if err == nil {
		v *= 2024
		newInscription = []string{ReduceZeros(fmt.Sprintf("%d", v))}
	}
	return newInscription
}
