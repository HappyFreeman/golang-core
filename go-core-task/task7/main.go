package main

import (
	"fmt"
	"time"
)

func main() {
	wg := NewCustomWaitGroup()

	for i := 1; i <= 5; i++ {
		wg.Add(1) // Добавляем 5 задач
		go func(i int) {
			defer wg.Done()
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("Task %d done\n", i)
		}(i)
	}

	wg.Wait()
	fmt.Println("All tasks are done!")
}
