package main

import (
	"fmt"
)

func main() {
	inputChan := make(chan uint8)
	outputChan := make(chan float64)

	// Запускаем конвейер
	createPipeline(inputChan, outputChan)

	// Пишем в первый канал числа
	go func() {
		for _, num := range []uint8{2, 3, 4, 5} {
			inputChan <- num
		}
		close(inputChan) // Закрываем канал после записи всех значений
	}()

	// Читаем результаты из второго канала
	for result := range outputChan {
		fmt.Printf("Result: %.2f\n", result)
	}
}
