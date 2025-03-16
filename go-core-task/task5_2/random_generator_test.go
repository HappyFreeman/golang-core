package main

import (
	"testing"
)

func TestGenerateRandomNumbers(t *testing.T) {
	count := 5
	ch := GenerateRandomNumbers(count)

	results := []int{}
	for num := range ch {
		results = append(results, num)
	}

	if len(results) != count {
		t.Errorf("Expected %d numbers, but got %d", count, len(results))
	}
}