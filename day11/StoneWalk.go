package main

import (
	"fmt"
	"log/slog"
	"sync"
	"sync/atomic"
)


func WalkNgen(firstGen []string, genCound int) int64 {
	total := int64(0)
	var waitGroup sync.WaitGroup
	for _, stone := range firstGen {
		waitGroup.Add(1)
		go func() {
			defer waitGroup.Done()
			sum := int64(1)
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
}
