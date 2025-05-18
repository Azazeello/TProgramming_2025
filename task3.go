package main

import (
	"fmt"
	"math"
)

func calculateY(a, b, x float64) float64 {
	return math.Sqrt((a + b*x) / math.Pow(math.Log10(x), 3))
}

func main() {
	// Задача A
	a, b := 7.2, 1.3
	xStart, xEnd, xStep := 1.56, 4.71, 0.63

	fmt.Println("Задача A:")
	for x := xStart; x <= xEnd; x += xStep {
		y := calculateY(a, b, x)
		fmt.Printf("x = %.2f, y = %.4f\n", x, y)
	}

	// Задача B
	xValues := []float64{2.4, 2.8, 3.9, 4.7, 3.16}

	fmt.Println("\nЗадача B:")
	for _, x := range xValues {
		y := calculateY(a, b, x)
		fmt.Printf("x = %.2f, y = %.4f\n", x, y)
	}
}