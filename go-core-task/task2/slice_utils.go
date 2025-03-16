package main

import "math/rand"

// Генерация случайного слайса
func generateRandomSlice(size int) []int {
	slice := make([]int, size)
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(100) // Генерация чисел от 0 до 99
	}
	return slice
}

func sliceExample(islice []int) []int {
	var oslice []int 

	for _, v := range islice {
		if v % 2 == 0 {
			oslice = append(oslice, v) // Добавляем четные числа в новый слайс
		}
	}
	return oslice
}

func addElement(slice []int, value int) []int {
	// append([]int{}, append(slice, value)...)
	return append(slice, value)
}

func copySlice(slice []int) []int {
	return append([]int{}, slice...)
}

// Удаляет элемент по индексу
func removeElement(slice []int, index int) []int {
	if index < 0 || index >= len(slice) {
        return slice
    }
	return append(slice[:index], slice[index+1:]...)
}
