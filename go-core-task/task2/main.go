package main

import (
	"fmt"
)

func main() {
	// Создаем слайс из 10 случайных значений
	originalSlice := generateRandomSlice(10)
	fmt.Println("Original Slice:", originalSlice)

	// Получаем только четные числа
	evenSlice := sliceExample(originalSlice)
	fmt.Println("Even Slice:", evenSlice)

	// Добавляем элемент в слайс
	newSlice := addElement(originalSlice, 42)
	fmt.Println("Slice after adding element:", newSlice)

	// Копируем слайс
	copiedSlice := copySlice(originalSlice)
	fmt.Println("Copied Slice:", copiedSlice)

	removedSlice := removeElement(originalSlice, 3)
	fmt.Println("Slice after removing element at index 3:", removedSlice)
}