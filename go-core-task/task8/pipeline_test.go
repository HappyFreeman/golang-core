package main

import (
	"testing"
)

// ok      command-line-arguments  0.465s

func TestPipeline(t *testing.T) {
	inputChan := make(chan uint8)
	outputChan := make(chan float64)

	// Запускаем конвейер
	createPipeline(inputChan, outputChan)

	go func() {
		inputChan <- 2
		inputChan <- 3
		inputChan <- 4
		close(inputChan)
	}()

	expectedResults := []float64{8, 27, 64}
	index := 0

	for result := range outputChan {
		if result != expectedResults[index] {
			t.Errorf("Expected %.2f, got %.2f", expectedResults[index], result)
		}
		index++
	}

	if index != len(expectedResults) {
		t.Errorf("Expected %d results, got %d", len(expectedResults), index)
	}
}
