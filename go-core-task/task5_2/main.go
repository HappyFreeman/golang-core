package main

import "fmt"

func main() {
	count := 10
	fmt.Printf("Generating %d random numbers:\n", count)

	// не блокируется, потому что это не операция записи и не операция чтения - это просто возврат канала 
	ch := GenerateRandomNumbers(count)

	for num := range ch {
		fmt.Println(num)
	}
}
