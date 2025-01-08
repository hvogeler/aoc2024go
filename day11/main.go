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
    flag.Parse()

	runtime.GOMAXPROCS(10)

	bytes, err := os.ReadFile("testdata.dat")
	if err != nil {
		panic(err)
	}
	raw := string(bytes)
	var data string = string(raw)
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

func Walk(stone string, sumStones *int64, currentGen int, maxGen int) {
	if currentGen >= maxGen {
		// slog.Info("     -- EXIT 1 --")
		return
	}
	// slog.Info(fmt.Sprintf("Walk Stone [%s] at Generation %d. Number of stones = %d\n", stone, currentGen, *sumStones))

	gen := currentGen
	stones := []string{stone}

	for len(stones) == 1 {
		stones = NextGen(stones)
		gen++
		if gen > maxGen {
			// slog.Info("     -- EXIT 2 --")
			return
		}
		// slog.Info(fmt.Sprintf("  Generation %d. Stones = %v\n", gen, stones))
	}

	if len(stones) == 2 {
		*sumStones++
		Walk(stones[0], sumStones, gen, maxGen)
		Walk(stones[1], sumStones, gen, maxGen)
		return
	}

	// panic(fmt.Sprintf("Wrong numger of stones for Walk. Can only be 1 or 2, but got %d", len(stones)))
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

func DoNgen(firstGen []string, n int) int64 {
	currentGen := firstGen
	for i := 0; i < n; i++ {
		nextGen := NextGen(currentGen)
		fmt.Printf("%d Gen %d stones", i+1, len(nextGen))
		if len(nextGen) < 20 {
			fmt.Printf(" %v\n", nextGen)
		} else {
			fmt.Println()
		}
		currentGen = nextGen
	}

	return int64(len(currentGen))
}
