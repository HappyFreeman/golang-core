package main

import "fmt"

func main() {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}

	// Проверяем пересечения
	hasIntersection, intersection := FindIntersection(a, b)

	fmt.Println("Slice A:", a)
	fmt.Println("Slice B:", b)
	fmt.Println("Has intersection:", hasIntersection)
	fmt.Println("Intersection:", intersection)
}