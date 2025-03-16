package main

import (
	"testing"
)

// ok      command-line-arguments  0.410s

func TestMergeChannels(t *testing.T) {
	ch1 := make(chan int, 2)
	ch2 := make(chan int, 2)
	ch3 := make(chan int, 2)

	ch1 <- 1
	ch1 <- 2
	close(ch1)

	ch2 <- 4
	ch2 <- 5
	close(ch2)

	ch3 <- 7
	ch3 <- 8
	close(ch3)

	merged := MergeChannels(ch1, ch2, ch3)

	results := []int{}
	for val := range merged {
		results = append(results, val)
	}

	if len(results) != 6 {
		t.Errorf("Expected 9 elements, but got %d", len(results))
	}
}
