package main

import (
	"fmt"
	"strings"
)

var scoringTable = map[string]int{
	"A X": 1 + 3,
	"A Y": 2 + 6,
	"A Z": 3 + 0,
	"B X": 1 + 0,
	"B Y": 2 + 3,
	"B Z": 3 + 6,
	"C X": 1 + 6,
	"C Y": 2 + 0,
	"C Z": 3 + 3,
}

func GetScore(strategy []string) int {
	score := 0
	for _, v := range strategy {
		roundScore := scoringTable[v]
		// TODO zero check to guard against bad input
		score += roundScore
	}
	return score
}

func main() {
	lines := strings.Split(puzzle, "\n")
	score := GetScore(lines)

	fmt.Printf("Part 1: Score %d\n", score)

}
