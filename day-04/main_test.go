package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRangeContainment(t *testing.T) {
	var cases = []struct {
		in  RangePair
		out bool
	}{
		{RangePair{Range{2, 4}, Range{6, 8}}, false},
		{RangePair{Range{2, 3}, Range{4, 5}}, false},
		{RangePair{Range{5, 7}, Range{7, 9}}, false},
		{RangePair{Range{4, 6}, Range{6, 6}}, true},
		{RangePair{Range{4, 6}, Range{5, 6}}, true},
		{RangePair{Range{4, 6}, Range{4, 6}}, true},
	}

	for _, tt := range cases {
		t.Run(fmt.Sprintf("Testing %v", tt.in), func(t *testing.T) {
			actual := tt.in.a.contains(tt.in.b)
			if tt.out != actual {
				t.Fatalf(`(%+v).contains(%+v) returned %t, expected %t`, tt.in.a, tt.in.b, actual, tt.out)
			}
		})
	}
}

func TestNewRange(t *testing.T) {
	var cases = []struct {
		in  string
		out Range
	}{
		{"2-4", Range{2, 4}},
		{"12-12", Range{12, 12}},
		{"0-9", Range{0, 9}},
	}
	for _, tt := range cases {
		t.Run(fmt.Sprintf("Testing %v", tt.in), func(t *testing.T) {
			actual, err := NewRange(tt.in)
			if err != nil || !reflect.DeepEqual(actual, &tt.out) {
				t.Fatalf(`NewRange(%s) returned %v, expected %v %s`, tt.in, actual, tt.out, err)
			}
		})
	}
}

func TestNewRangePair(t *testing.T) {
	var cases = []struct {
		in  string
		out RangePair
	}{
		{"2-4,6-8", RangePair{Range{2, 4}, Range{6, 8}}},
		{"2-3,4-5", RangePair{Range{2, 3}, Range{4, 5}}},
		{"0-99,78-88", RangePair{Range{0, 99}, Range{78, 88}}},
	}

	for _, tt := range cases {
		t.Run(fmt.Sprintf("Testing %v", tt.in), func(t *testing.T) {
			actual, err := NewRangePair(tt.in)
			if err != nil || !reflect.DeepEqual(actual, &tt.out) {
				t.Fatalf(`NewRangePair(%s) returned %v, expected %v %s`, tt.in, actual, tt.out, err)
			}
		})
	}
}

func TestCountContained(t *testing.T) {
	lines := []string{
		"2-4,6-8",
		"2-3,4-5",
		"5-7,7-9",
		"2-8,3-7",
		"6-6,4-6",
		"2-6,4-8",
	}
	rangePairs, err := LinesToRangePairs(lines)
	if err != nil {
		t.Fatalf("Clould not convert %v to range pairs: %s", lines, err)
	}

	expected := 2
	actual := CountContained(rangePairs)
	if actual != expected {
		t.Fatalf(`CountContained(%v) returned %d, expected %d`, rangePairs, actual, expected)
	}
}
