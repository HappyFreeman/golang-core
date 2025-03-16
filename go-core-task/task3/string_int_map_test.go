package main

import (
	"reflect"
	"testing"
)

// ok      command-line-arguments  0.415s

func TestAddAndGet(t *testing.T) {
	m := StringIntMap{
		data: make(map[string]int),
	}

	m.Add("apple", 10)

	value, exists := m.Get("apple")
	if !exists || value != 10 {
		t.Errorf("Expected value 10, got %d", value)
	}

	_, exists = m.Get("banana")
	if exists {
		t.Errorf("Expected 'banana' not to exist")
	}
}

func TestRemove(t *testing.T) {
	m := StringIntMap{
		data: make(map[string]int),
	}

	m.Add("apple", 10)
	m.Remove("apple")

	_, exists := m.Get("apple")
	if exists {
		t.Errorf("Expected 'apple' to be removed")
	}
}

func TestCopy(t *testing.T) {
	m := StringIntMap{
		data: make(map[string]int),
	}

	m.Add("apple", 10)
	m.Add("banana", 20)

	copyMap := m.Copy()
	if !reflect.DeepEqual(m, copyMap) {
		t.Errorf("Expected %v, got %v", m, copyMap)
	}

	// Изменяем копию и проверяем, что оригинал не изменился
	copyMap.Add("apple", 100)
	value, _ := m.Get("apple")
	if value == 100 {
		t.Errorf("Expected original value to remain unchanged")
	}
}