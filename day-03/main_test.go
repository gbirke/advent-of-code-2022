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

var testData = []struct {
	in  string
	out rune
}{
	{"vJrwpWtwJgWrhcsFMMfFFhFp", 'p'},
	{"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", 'L'},
	{"PmmdzqPrVvPwwTWBwg", 'P'},
	{"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn", 'v'},
	{"ttgJtRGJQctTZtZT", 't'},
	{"CrZsJsPPZsGzwwsLwLmpwMDw", 's'},
}

func TestFindDuplicate(t *testing.T) {
	for _, tt := range testData {
		t.Run(fmt.Sprintf("Testing %s", tt.in), func(t *testing.T) {
			actual, err := FindDuplicate(tt.in)
			if err != nil || tt.out != actual {
				t.Fatalf(`FindDuplicate(%s) returned %c, expected %c %s`, tt.in, actual, tt.out, err)
			}
		})
	}
}

func TestCalculayePriority(t *testing.T) {
	var lines []string
	for _, td := range testData {
		lines = append(lines, td.in)
	}
	expected := 157
	actual, err := CalculatePriority(lines)
	if err != nil || expected != actual {
		t.Fatalf(`CalculatePriority(%v) returned %d, expected %d %s`, lines, actual, expected, err)
	}
}
