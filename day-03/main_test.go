package main

import (
	"fmt"
	"testing"
)

func TestCharacterToPriority(t *testing.T) {
	var cases = []struct {
		in  rune
		out int
	}{
		{'a', 1},
		{'b', 2},
		{'c', 3},
		{'z', 26},
		{'A', 27},
		{'B', 28},
		{'Z', 52},
	}

	for _, tt := range cases {
		t.Run(fmt.Sprintf("Testing %c", tt.in), func(t *testing.T) {
			actual := CharacterToPriority(tt.in)
			if tt.out != actual {
				t.Fatalf(`CharacterToPriority(%c) returned %b, expected %b`, tt.in, actual, tt.out)
			}
		})
	}
}

var testInput = []string{
	"vJrwpWtwJgWrhcsFMMfFFhFp",
	"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
	"PmmdzqPrVvPwwTWBwg",
	"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
	"ttgJtRGJQctTZtZT",
	"CrZsJsPPZsGzwwsLwLmpwMDw",
}

func TestFindDuplicate(t *testing.T) {
	var expectedChars = []rune{
		'p',
		'L',
		'P',
		'v',
		't',
		's',
	}

	for idx, tt := range testInput {
		t.Run(fmt.Sprintf("Testing %s", tt), func(t *testing.T) {
			actual, err := FindDuplicate(tt)
			if err != nil || expectedChars[idx] != actual {
				t.Fatalf(`FindDuplicate(%s) returned %c, expected %c %s`, tt, actual, expectedChars[idx], err)
			}
		})
	}
}

func TestFindCommonChar(t *testing.T) {
	expectedChars := map[int]rune{0: 'r', 3: 'Z'}
	for i := 0; i < len(testInput); i += 3 {
		comparisonSlice := testInput[i : i+3]
		t.Run(fmt.Sprintf("Testing %v", comparisonSlice), func(t *testing.T) {
			actual, err := FindCommonChar(comparisonSlice)
			if err != nil || expectedChars[i] != actual {
				t.Fatalf(`TestFindCommonChar(%v) returned %c, expected %c %s`, comparisonSlice, actual, expectedChars[i], err)
			}
		})
	}
}

func TestCalculatePriority(t *testing.T) {
	expected := 157
	actual, err := CalculatePriority(testInput)
	if err != nil || expected != actual {
		t.Fatalf(`CalculatePriority(%v) returned %d, expected %d %s`, testInput, actual, expected, err)
	}
}

func TestCalculatePriorityByShared(t *testing.T) {
	expected := 70
	actual, err := CalculatePriorityByShared(testInput)
	if err != nil || expected != actual {
		t.Fatalf(`CalculatePriorityByShared(%v) returned %d, expected %d %s`, testInput, actual, expected, err)
	}
}
