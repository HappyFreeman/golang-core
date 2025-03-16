package main

import (
	"reflect"
	"testing"
)

// ok      command-line-arguments  0.396s

func TestFindIntersection(t *testing.T) {
	tests := []struct {
		a           []int
		b           []int
		hasIntersect bool
		intersection []int
	}{
		{
			a:            []int{65, 3, 58, 678, 64},
			b:            []int{64, 2, 3, 43},
			hasIntersect:  true,
			intersection: []int{64, 3},
		},
		{
			a:            []int{1, 2, 3},
			b:            []int{4, 5, 6},
			hasIntersect:  false,
			intersection: []int{},
		},
		{
			a:            []int{1, 2, 3},
			b:            []int{},
			hasIntersect:  false,
			intersection: []int{},
		},
		{
			a:            []int{},
			b:            []int{1, 2, 3},
			hasIntersect:  false,
			intersection: []int{},
		},
		{
			a:            []int{1, 2, 3, 4},
			b:            []int{2, 3, 4, 5},
			hasIntersect:  true,
			intersection: []int{2, 3, 4},
		},
	}

	for _, test := range tests {
		hasIntersect, intersection := FindIntersection(test.a, test.b)

		if hasIntersect != test.hasIntersect || !reflect.DeepEqual(intersection, test.intersection) {
            t.Errorf("Expected %t for a=%v and b=%v, got %t", test.hasIntersect, test.a, test.b, hasIntersect)
        }
	}
}