package main

import (
	"fmt"
	"strconv"
	"strings"
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

func main() {
	lines := strings.Split(puzzle, "\n")
	calorieElves := BuildMap(lines)
	maxElf := GetBiggestKeyFromMap(calorieElves)
	fmt.Printf("Part 1: Elf number %v\n", calorieElves[maxElf])

}
