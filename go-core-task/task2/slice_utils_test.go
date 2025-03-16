package main

import (
	"reflect"
	"testing"
)

// ok      command-line-arguments  0.394s

func TestSliceExample(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6}
	expected := []int{2, 4, 6}
	result := sliceExample(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestAddElement(t *testing.T) {
	input := []int{1, 2, 3}
	result := addElement(input, 4)
	expected := []int{1, 2, 3, 4}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestCopySlice(t *testing.T) {
	input := []int{1, 2, 3}
	result := copySlice(input)
	if !reflect.DeepEqual(result, input) {
		t.Errorf("Expected %v, got %v", input, result)
	}

	// Проверяем, что изменения в копии не влияют на оригинал
	result[0] = 42
	if reflect.DeepEqual(result, input) {
		t.Errorf("Changes in the copy affected the original slice")
	}
}

func TestRemoveElement(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	result := removeElement(input, 2)
	expected := []int{1, 2, 4, 5}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}