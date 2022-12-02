package main

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/gbirke/advent-of-code-2022/pkg"
)

func BuildMap(lines []string) map[int]int {
	m := make(map[int]int)
	index := 0
	for _, line := range lines {
		if line != "" {
			lineValue, err := strconv.Atoi(line)
			if err != nil {
				fmt.Printf("Cannot parse %s\n", line)
				// TODO return err instead
				continue
			}
			m[index] += lineValue
		} else {
			index += 1
		}
	}
	return m
}

func GetBiggestKeyFromMap(inputMap map[int]int) int {
	var result, currentMax int
	for k, v := range inputMap {
		if v > currentMax {
			currentMax = v
			result = k
		}
	}
	return result
}

func GetSortedValuesFromMap(inputMap map[int]int) []int {
	values := make([]int, 0, len(inputMap))
	for _, v := range inputMap {
		values = append(values, v)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(values)))
	return values
}

func main() {
	lines, err := pkg.ReadPuzzle()
	if err != nil {
		panic(err)
	}

	calorieElves := BuildMap(lines)
	sortedCalories := GetSortedValuesFromMap(calorieElves)
	fmt.Printf("Part 1: Most calories %v\n", sortedCalories[0])

	max3 := 0

	for i := 0; i < 3; i++ {
		max3 += sortedCalories[i]
	}
	fmt.Printf("Part 2: Most 3 calories %v\n", max3)

}
