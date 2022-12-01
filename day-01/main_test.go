package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBuildMapWithIntegerList(t *testing.T) {
	lines := []string{"111", "222", "333"}
	expected := map[int]int{0: 666}
	builtMap := BuildMap(lines)
	if !reflect.DeepEqual(builtMap, expected) {
		t.Fatalf(`BuildMap(%v) returned %+v, expected %+v`, lines, builtMap, expected)
	}
}

func TestBuildMapWithSpacesIntegerList(t *testing.T) {
	lines := []string{"111", "", "222", "333", "", "7"}
	expected := map[int]int{0: 111, 1: 555, 2: 7}
	builtMap := BuildMap(lines)
	if !reflect.DeepEqual(builtMap, expected) {
		t.Fatalf(`BuildMap(%v) returned %+v, expected %+v`, lines, builtMap, expected)
	}
}

var keyFromMapTests = []struct {
	in  map[int]int
	out int
}{
	{map[int]int{0: 111, 1: 555, 2: 7}, 1},
	{map[int]int{0: 666, 1: 555, 2: 7}, 0},
	{map[int]int{0: 0, 1: 0, 2: 7}, 2},
}

func TestGetBiggestKeyFromMap(t *testing.T) {
	for _, tt := range keyFromMapTests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			i := GetBiggestKeyFromMap(tt.in)
			if i != tt.out {
				t.Errorf("GetBiggestKeyFromMap(%v) => %v, want %v", tt.in, i, tt.out)
			}
		})
	}

}
