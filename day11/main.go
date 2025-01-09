package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

func main() {

	// stone := "125"
	// genCound := 50
	// sum := 1

	// slog.Info("Calculate %d generations for stone %s", genCound, stone)
	// Walk(stone, &sum, 0, genCound)
	// slog.Info(fmt.Sprintf("Stone %s after %d generations. Number of stones: %d\n", stone, genCound, sum))

	ngen := flag.Int("ngen", 25, "number of generations to calculate")
	walk := flag.Bool("walk", true, "do recursive walk (true) or loop (false)")
	pdata := flag.String("data", "", "Stones like '125 17 3'")
	flag.Parse()

	runtime.GOMAXPROCS(10)
	var data string
	if *pdata == "" {
		bytes, err := os.ReadFile("testdata.dat")
		if err != nil {
			panic(err)
		}
		raw := string(bytes)
		data = string(raw)
	} else {
		data = *pdata
	}
	firstGen := strings.Split(data, " ")

	// firstGen := []string{"872027"} // took 9:20
	start := time.Now()
	var n int64
	if *walk {
		n = WalkNgen(firstGen, *ngen)
	} else {
		n = DoNgen(firstGen, *ngen)
	}
	duration := time.Since(start)
	slog.Info(fmt.Sprintf("Stones: %d  ---  run duration: %v", n, duration))
	// // Part 1: 199946

	// sum := 0
	// for i, stone := range firstGen {
	// 	fmt.Printf("Stone %s, %d of %d\n", stone, i+1, len(firstGen))
	// 	n = DoNgen([]string{stone}, 75)
	// 	fmt.Printf("Stone %s, %d of %d, Stones %d", stone, i+i, len(firstGen), n)
	// 	sum += n
	// }
	// fmt.Println("Stones: ", sum)

}

func WalkNgen(firstGen []string, genCound int) int64 {
	total := int64(0)
	var waitGroup sync.WaitGroup
	for _, stone := range firstGen {
		waitGroup.Add(1)
		go func() {
			defer waitGroup.Done()
			sum := int64(1)
			// slog.Info("\n\n-----------------------------------------------------------\n")
			slog.Info(fmt.Sprintf("Calculate %d generations for stone %s", genCound, stone))
			Walk(stone, &sum, 0, genCound)
			slog.Info(fmt.Sprintf("Done. Stone %s after %d generations. Number of stones: %d\n", stone, genCound, sum))
			atomic.AddInt64(&total, sum)
		}()
	}
	waitGroup.Wait()
	return total
}

var paths = 0

func Walk(stone string, sumStones *int64, currentGen int, maxGen int) {
	paths++
	if currentGen >= maxGen {
		slog.Info("    -- EXIT 1 --", "Gen", currentGen, "Depth", paths)
		paths--
		return
	}
	// slog.Info(fmt.Sprintf("Walk Stone [%s] at Generation %d. Number of stones = %d\n", stone, currentGen, *sumStones))

	gen := currentGen
	stones := []string{stone}

	for len(stones) == 1 {
		stones = NextGen(stones)
		gen++
		if gen > maxGen {
			paths--
			slog.Info("     -- EXIT 2 --", "Gen", gen, "Depth", paths)
			return
		}
		// slog.Info(fmt.Sprintf("  Generation %d. Stones = %v\n", gen, stones))
	}

	if len(stones) == 2 {
		*sumStones++
		Walk(stones[0], sumStones, gen, maxGen)
		Walk(stones[1], sumStones, gen, maxGen)
		paths--
		slog.Info("     -- EXIT 3 --", "Gen", gen, "Depth", paths)
		return
	}

	// panic(fmt.Sprintf("Wrong numger of stones for Walk. Can only be 1 or 2, but got %d", len(stones)))
}

type StoneMap map[string]int

func NextGenMap(currentGen *StoneMap) StoneMap {
	nextGen := make(StoneMap)
	for inscription, count := range *currentGen {
		fmt.Println(inscription, count)
		newInscription := applyRules(inscription)
		switch len(newInscription) {
		case 1:
			nextGen[newInscription[0]] = count
        case 2:
			nextGen[newInscription[0]] = count
			nextGen[newInscription[1]] = count
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

func DoNgenMap(firstGen StoneMap, n int) int64 {
	currentGen := firstGen
	for i := 0; i < n; i++ {
		nextGen := NextGenMap(&currentGen)
		// fmt.Printf("%d\n", len(nextGen))
		if len(nextGen) < 150 {
			fmt.Printf(" %v\n", nextGen)
		} else {
			fmt.Println()
		}
		currentGen = nextGen
	}

	return int64(CountStones(currentGen))
}

func CountStones(stones StoneMap) int64 {
    sum := int64(0)
    for _, count := range stones {
        sum += int64(count)
    }
    return sum
}

func DoNgen(firstGen []string, n int) int64 {
	currentGen := firstGen
	for i := 0; i < n; i++ {
		nextGen := NextGen(currentGen)
		fmt.Printf("%d Gen %d stones, stones %d", i+1, len(nextGen), len(nextGen))
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

func dedup(stones *[]string) StoneMap {
	m := make(map[string]int)
	for _, stone := range *stones {
		m[stone] = m[stone] + 1
	}
	return m
}
