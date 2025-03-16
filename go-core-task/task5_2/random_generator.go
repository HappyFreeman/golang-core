package main

import (
	"math/rand"
)

// GenerateRandomNumbers генерирует случайные числа и отправляет их в небуферизированный канал.
func GenerateRandomNumbers(count int) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)
		for i := 0; i < count; i++ {
			ch <- rand.Intn(100) // Генерируем числа от 0 до 99
		}
	}()

	return ch
}
