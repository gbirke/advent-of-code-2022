package main

import (
	"testing"
)

func TestGetScorePart1(t *testing.T) {
	lines := []string{"A Y", "B X", "C Z"}
	want := 15
	actual := GetScore(lines, naiveScoringTable)
	if want != actual {
		t.Fatalf(`GetScore(%v) => %d, want %d`, lines, actual, want)
	}
}

func TestGetScorePart2(t *testing.T) {
	lines := []string{"A Y", "B X", "C Z"}
	want := 12
	actual := GetScore(lines, cleverScoringTable)
	if want != actual {
		t.Fatalf(`GetScore(%v) => %d, want %d`, lines, actual, want)
	}
}
