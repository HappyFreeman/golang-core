package main

import (
	"fmt"
)

func main() {
	m := StringIntMap{
		data: make(map[string]int),
	}

	m.Add("apple", 5)
	m.Add("banana", 3)

	fmt.Println(m)

	// Получаем значение по ключу
	value, exists := m.Get("banana")
	if exists {
		fmt.Println("Value for 'banana':", value)
	} else {
		fmt.Println("'banana' not found")
	}

	fmt.Println("Exists 'apple':", m.Exists("apple"))
	fmt.Println("Exists 'apple':", m.Exists("sda"))

	copyMap := m.Copy()
	m.Remove("apple")

	fmt.Println("Original map:", m)
	fmt.Println("copy map:", copyMap)
}