package main

import (
	"testing"
)

func TestGetScore(t *testing.T) {
	lines := []string{"A Y", "B X", "C Z"}
	want := 15
	actual := GetScore(lines)
	if want != actual {
		t.Fatalf(`GetScore(%v) => %d, want %d`, lines, actual, want)
	}
}
