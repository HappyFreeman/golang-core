package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	numDecimal := 42           // Десятичная система
	numOctal := 052            // Восьмеричная система
	numHexadecimal := 0x2A     // Шестнадцатиричная система
	pi := 3.14                 // Тип float64
	name := "Golang"           // Тип string
	isActive := true           // Тип bool
	complexNum := complex64(complex(1, 2)) // Тип complex64

	// Определяем типы переменных и выводим их на экран
	fmt.Printf("numDecimal: %T = %v\n", numDecimal, numDecimal)
	fmt.Printf("numOctal: %T = %v\n", numOctal, numOctal)
	fmt.Printf("numHexadecimal: %T = %v\n", numHexadecimal, numHexadecimal)
	fmt.Printf("pi: %T = %v\n", pi, pi)
	fmt.Printf("name: %T = %v\n", name, name)
	fmt.Printf("isActive: %T = %v\n", isActive, isActive)
	fmt.Printf("complexNum: %T = %v\n", complexNum, complexNum)

	// Преобразует все переменные в строковый тип и объединяет их в одну строку.
	combined := fmt.Sprintf("%d %d %d %.3f %s %t %v", numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum)
	fmt.Printf("combined: %s\n", combined)

	// Преобразуем строку в срез рун
	runes := []rune(combined)
	fmt.Println("Rune slice:", runes)

	// Добавляем соль в середину строки и хэшируем результат
	mid := len(runes) / 2
	runes = append(runes[:mid], append([]rune("go-2024"), runes[mid:]...)...)
	hash := sha256.Sum256([]byte(string(runes)))
	fmt.Println("SHA256 Hash with salt:", hash)
}