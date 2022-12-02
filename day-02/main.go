package main

import (
	"fmt"

	"github.com/gbirke/advent-of-code-2022/pkg"
)

const Rock = 1
const Paper = 2
const Scissors = 3

var naiveScoringTable = map[string]int{
	"A X": Rock + 3,
	"A Y": Paper + 6,
	"A Z": Scissors + 0,
	"B X": Rock + 0,
	"B Y": Paper + 3,
	"B Z": Scissors + 6,
	"C X": Rock + 6,
	"C Y": Paper + 0,
	"C Z": Scissors + 3,
}

var cleverScoringTable = map[string]int{
	"A X": Scissors + 0,
	"A Y": Rock + 3,
	"A Z": Paper + 6,
	"B X": Rock + 0,
	"B Y": Paper + 3,
	"B Z": Scissors + 6,
	"C X": Paper + 0,
	"C Y": Scissors + 3,
	"C Z": Rock + 6,
}

func GetScore(strategy []string, scoringTable map[string]int) int {
	score := 0
	for _, v := range strategy {
		roundScore := scoringTable[v]
		// TODO zero check to guard against bad input
		score += roundScore
	}
	return score
}

func main() {
	lines, err := pkg.ReadPuzzle()
	if err != nil {
		panic(err)
	}

	score := GetScore(lines, naiveScoringTable)

	fmt.Printf("Part 1: Score %d\n", score)

	score2 := GetScore(lines, cleverScoringTable)

	fmt.Printf("Part 2: Score %d\n", score2)

}
