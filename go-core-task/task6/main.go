package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	// Заполняем каналы через горутины
	go func() {
		for i := 1; i <= 3; i++ {
			ch1 <- i
			time.Sleep(200 * time.Millisecond)
		}
		close(ch1)
	}()

	go func() {
		for i := 4; i <= 6; i++ {
			ch2 <- i
			time.Sleep(100 * time.Millisecond)
		}
		close(ch2)
	}()

	go func() {
		for i := 7; i <= 9; i++ {
			ch3 <- i
			time.Sleep(150 * time.Millisecond)
		}
		close(ch3)
	}()

	// Слияние каналов
	merged := MergeChannels(ch1, ch2, ch3)

	// Читаем данные из объединённого канала
	for val := range merged {
		fmt.Println(val)
	}
}
