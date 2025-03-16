package main

import (
	"reflect"
	"testing"
)

// ok      command-line-arguments  0.443s

func TestDifference(t *testing.T) {
	tests := []struct {
		slice1   []string
		slice2   []string
		expected []string
	}{
		{
			slice1:   []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"},
			slice2:   []string{"banana", "date", "fig"},
			expected: []string{"apple", "cherry", "43", "lead", "gno1"},
		},
		{
			slice1:   []string{"apple", "banana", "cherry"},
			slice2:   []string{"apple", "banana", "cherry"},
			expected: []string{},
		},
		{
			slice1:   []string{"apple", "banana", "cherry"},
			slice2:   []string{},
			expected: []string{"apple", "banana", "cherry"},
		},
		{
			slice1:   []string{},
			slice2:   []string{"apple", "banana", "cherry"},
			expected: []string{},
		},
	}

	for _, test := range tests {
		result := Difference(test.slice1, test.slice2)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Difference(%v, %v) = %v; want %v", test.slice1, test.slice2, result, test.expected)
		}
	}
}
