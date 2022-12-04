package main

import (
	"errors"
	"fmt"

	"github.com/gbirke/advent-of-code-2022/pkg"
)

const asciiMask = 0b11100000

func CharacterToPriority(c rune) int {
	firstBytes := int(c)

	// Look at bit 5 to see if char is uppercase or lowercase ASCI
	caseBit := ((firstBytes >> 5) &^ 0b110)

	var offset int
	if caseBit == 0 {
		offset = 26
	}

	return (firstBytes &^ asciiMask) + offset
}

func FindDuplicate(str string) (rune, error) {
	halfIndex := len(str) / 2
	first := str[0:halfIndex]
	second := str[halfIndex:]

	return FindCommonChar([]string{first, second})
}

func FindCommonChar(str []string) (rune, error) {

	seen := make(map[rune]int)
	lastIndex := len(str) - 1
	var foundInAllPrevious int
	for i := 0; i < lastIndex; i++ {
		indexBit := 0b1 << i
		foundInAllPrevious = foundInAllPrevious | indexBit
		for _, c := range str[i] {
			seen[c] = seen[c] | indexBit
		}
	}

	for _, c := range str[lastIndex] {
		if seen[c] == foundInAllPrevious {
			return c, nil
		}
	}
	return '?', errors.New("Strings don't have a common character")
}

func CalculatePriority(lines []string) (int, error) {
	var result int
	for _, l := range lines {
		duplicate, err := FindDuplicate(l)
		if err != nil {
			return 0, err
		}
		result += CharacterToPriority(duplicate)
	}
	return result, nil
}

func main() {
	lines, err := pkg.ReadPuzzle()
	if err != nil {
		panic(err)
	}

	priority, err := CalculatePriority(lines)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part 1 Priority: %d\n", priority)
}
