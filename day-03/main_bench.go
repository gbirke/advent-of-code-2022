package main

import (
	"testing"

	"github.com/gbirke/advent-of-code-2022/pkg"
)

func BenchmarkSolution(b *testing.B) {
	lines, err := pkg.ReadPuzzle()
	if err != nil {
		panic(err)
	}

	_, err = CalculatePriority(lines)
	if err != nil {
		panic(err)
	}

	_, err = CalculatePriorityByShared(lines)
	if err != nil {
		panic(err)
	}

}
