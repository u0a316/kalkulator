package main

import (
	"fmt"
	"os"
	"strconv"
)

func add(a, b float64) float64 {
	return a + b
}

func multiply(a, b float64) float64 {
	return a * b
}

func main() {
	fmt.Println("Go Calculator")

	// Pastikan ada dua argumen yang diberikan
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run calc.go <number1> <number2>")
		return
	}

	// Konversi argumen menjadi float64
	num1, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println("Invalid number 1:", os.Args[1])
		return
	}

	num2, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil {
		fmt.Println("Invalid number 2:", os.Args[2])
		return
	}

	// Hitung hasil penambahan dan perkalian
	sum := add(num1, num2)
	product := multiply(num1, num2)

	// Tampilkan hasil
	fmt.Printf("Result of addition: %.2f\n", sum)
	fmt.Printf("Result of multiplication: %.2f\n", product)
}

