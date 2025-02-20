package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"day13/clawmachine"
)

func main() {
	bytes, err := os.ReadFile("testdata.dat")
	if err != nil {
		panic(err)
	}
	data := string(bytes)
	// fmt.Println(data)

	games := games(data)
	sumTokens := 0
	for i, game := range games {
		machine := clawmachine.MachinefromStr(game)
		totalPressesA, totalPressesB := machine.FindPrize()
		cost, err := clawmachine.LowestCost(totalPressesA, totalPressesB)
		if err == nil {
			fmt.Printf("Game %5d costs %d tokens\n", i+1, cost)
			sumTokens += cost
		} else {
			fmt.Printf("Game %5d did not gain a prize: %s\n", i+1, *err)
		}
	}
	fmt.Printf("Part1 game cost: %d tokens\n\n", sumTokens)
	// Part1 Result: 33481

	sumTokens = 0
	for i, game := range games {
		machine := clawmachine.MachinefromStr(game)
		machine.IncreasePrizeLocationForPart2(10000000000000)
		totalPressesA, totalPressesB, err := machine.FindPrize2()
		if err == nil {
            cost := clawmachine.Cost(totalPressesA, totalPressesB)
			fmt.Printf("Game %5d costs %d tokens\n", i+1, cost)
			sumTokens += cost
		} else {
			fmt.Printf("Game %5d did not gain a prize: %s\n", i+1, *err)
		}
	}
	fmt.Printf("Part2 game cost: %d tokens\n", sumTokens)


}

func games(s string) []string {
	games := []string{}
	game := ""
	scanner := bufio.NewScanner(strings.NewReader(s))
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			games = append(games, game)
			game = ""
		} else {
			game += (line + "\n")
		}
	}
	return games
}