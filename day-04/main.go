package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gbirke/advent-of-code-2022/pkg"
)

type Range struct {
	lower int
	upper int
}

type RangePair struct {
	a, b Range
}

func (r Range) contains(comparisonRange Range) bool {
	return r.lower <= comparisonRange.lower &&
		r.upper >= comparisonRange.upper
}

func (r Range) overlaps(comparisonRange Range) bool {
	return (r.lower <= comparisonRange.lower && r.upper >= comparisonRange.lower) ||
		(r.upper >= comparisonRange.upper && r.lower <= comparisonRange.upper)
}

func NewRange(str string) (*Range, error) {
	pairs := strings.SplitN(str, "-", 2)
	lower, err := strconv.Atoi(pairs[0])
	if err != nil {
		return nil, err
	}
	upper, err := strconv.Atoi(pairs[1])
	if err != nil {
		return nil, err
	}
	return &Range{lower, upper}, nil
}

func NewRangePair(str string) (*RangePair, error) {
	pairs := strings.SplitN(str, ",", 2)
	a, err := NewRange(pairs[0])
	if err != nil {
		return nil, err
	}
	b, err := NewRange(pairs[1])
	if err != nil {
		return nil, err
	}
	return &RangePair{*a, *b}, nil
}

func LinesToRangePairs(lines []string) ([]RangePair, error) {
	var rangePairs []RangePair
	for _, line := range lines {
		rangePair, err := NewRangePair(line)
		if err != nil {
			return rangePairs, err
		}
		rangePairs = append(rangePairs, *rangePair)
	}
	return rangePairs, nil
}

func CountContained(rangePairs []RangePair) int {
	var result int
	for _, rp := range rangePairs {
		if rp.a.contains(rp.b) || rp.b.contains(rp.a) {
			result += 1
		}
	}
	return result
}

func CountOverlaps(rangePairs []RangePair) int {
	var result int
	for _, rp := range rangePairs {
		if rp.a.overlaps(rp.b) || rp.b.overlaps(rp.a) {
			result += 1
		}
	}
	return result
}

func main() {
	lines, err := pkg.ReadPuzzle()
	if err != nil {
		panic(err)
	}
	rangePairs, err := LinesToRangePairs(lines)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1 contained ranges: %d\n", CountContained(rangePairs))
	fmt.Printf("Part 2 overlapping ranges: %d\n", CountOverlaps(rangePairs))
}
